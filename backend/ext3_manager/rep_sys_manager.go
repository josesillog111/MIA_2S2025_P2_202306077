package ext3

import (
	dmanager "backend/disk_manager"
	file "backend/ext3_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type ReportSysManager struct {
	Partition        dmanager.MountedPartition
	SBManager        *SBManager
	FileManager      FileManager
	InodeManager     InodeManager
	BlockManager     BlockManager
	DirectoryManager DirectoryManager
}

func NewReportSysManager(partition dmanager.MountedPartition, sb *SBManager, fm FileManager, im InodeManager, bm BlockManager, dm DirectoryManager) *ReportSysManager {
	return &ReportSysManager{
		Partition:        partition,
		SBManager:        sb,
		FileManager:      fm,
		InodeManager:     im,
		BlockManager:     bm,
		DirectoryManager: dm,
	}
}

func (r *ReportSysManager) Inode(path string, format string) error {
	var b strings.Builder

	b.WriteString("digraph G {\n")
	b.WriteString("rankdir=LR;\n")
	b.WriteString("node [shape=box];\n\n")

	sb := r.SBManager.SB
	clusterID := 0
	var nodes []string

	for inodeID := int64(0); inodeID < sb.S_inodes_count; inodeID++ {
		inode, err := r.FileManager.InodeManager.ReadInode(inodeID)
		if err != nil {
			// inodo no legible → saltar
			continue
		}

		// Filtrar por tipo (solo DIR o FILE)
		if inode.I_type != DIR_TYPE && inode.I_type != FILE_TYPE {
			continue
		}

		// Detección estricta de "inodo vacío":
		// Si no tiene UID, GID, size, permisos ni timestamps -> lo consideramos vacío
		if inode.I_uid == 0 &&
			inode.I_gid == 0 &&
			inode.I_size == 0 &&
			inode.I_permissions == 0 &&
			inode.I_atime == 0 &&
			inode.I_ctime == 0 &&
			inode.I_mtime == 0 {
			// inodo aparentemente no usado → saltar
			continue
		}

		// Tipo legible
		var inodeType string
		switch inode.I_type {
		case DIR_TYPE:
			inodeType = "DIR"
		case FILE_TYPE:
			inodeType = "FILE"
		default:
			inodeType = "UNK"
		}

		// Construir tabla con la información del inodo
		var rows []string
		// uid
		rows = append(rows, fmt.Sprintf("<TR><TD>i_uid</TD><TD>%d</TD></TR>", inode.I_uid))
		// gid
		rows = append(rows, fmt.Sprintf("<TR><TD>i_gid</TD><TD>%d</TD></TR>", inode.I_gid))
		//size
		rows = append(rows, fmt.Sprintf("<TR><TD>size</TD><TD>%d</TD></TR>", inode.I_size))
		// atime
		rows = append(rows, fmt.Sprintf("<TR><TD>i_atime</TD><TD>%s</TD></TR>", time.Unix(inode.I_atime, 0).Format("2006-01-02 15:04:05")))
		// ctime
		rows = append(rows, fmt.Sprintf("<TR><TD>i_ctime</TD><TD>%s</TD></TR>", time.Unix(inode.I_ctime, 0).Format("2006-01-02 15:04:05")))
		// mtime
		rows = append(rows, fmt.Sprintf("<TR><TD>i_mtime</TD><TD>%s</TD></TR>", time.Unix(inode.I_mtime, 0).Format("2006-01-02 15:04:05")))
		// type
		rows = append(rows, fmt.Sprintf("<TR><TD>i_type</TD><TD>%s</TD></TR>", inodeType))
		// permissions
		rows = append(rows, fmt.Sprintf("<TR><TD>i_permissions</TD><TD>%o</TD></TR>", inode.I_permissions))

		// Mostrar los 15 punteros (incluir solo los que no sean -1)
		for idx, blk := range inode.I_block {
			if blk != -1 {
				rows = append(rows, fmt.Sprintf("<TR><TD>Bloque[%d]</TD><TD>%d</TD></TR>", idx, blk))
			}
		}

		nodeName := fmt.Sprintf("inode_%d", inodeID)
		nodes = append(nodes, nodeName)

		b.WriteString(fmt.Sprintf("    subgraph cluster_%d {\n", clusterID))
		b.WriteString(fmt.Sprintf("        label = \"Inodo %d\";\n", inodeID))
		b.WriteString("        color=white;\n        node [shape=box, margin=0.2, fillcolor=white];\n\n")
		b.WriteString(fmt.Sprintf("        \"%s\" [label=<<TABLE BORDER=\"0\" CELLBORDER=\"0\" CELLSPACING=\"0\">\n", nodeName))
		for _, row := range rows {
			b.WriteString("            " + row + "\n")
		}
		b.WriteString("        </TABLE>>];\n    }\n\n")
		clusterID++
	}

	// Conectar inodos activos como lista enlazada (orden de aparición)
	for i := 0; i < len(nodes)-1; i++ {
		b.WriteString(fmt.Sprintf("    %s -> %s;\n", nodes[i], nodes[i+1]))
	}

	b.WriteString("}\n")
	// Generar gráfico (usa render si format != "dot" dentro de graphvizExecute)
	const generate_dot bool = false
	err := r.graphvizExecute(b, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("Inode: Error generando gráfico: %v", err)
	}

	return nil
}

func (r *ReportSysManager) Block(path string, format string) error {
	sb, err := r.SBManager.ReadSuperBlock()
	if err != nil {
		return fmt.Errorf("Block: Error leyendo superbloque: %v", err)
	}

	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("    rankdir=LR;\n")
	b.WriteString("    node [shape=box];\n\n")

	clusterID := 0
	var nodes []string // nombres de nodos principales

	for i := int64(0); i < sb.S_blocks_count; i++ {
		used, err := r.BlockManager.IsBlockUsed(i)
		if err != nil {
			return fmt.Errorf("Block: Error leyendo bitmap de bloque %d: %v", i, err)
		}
		if !used {
			continue
		}

		raw, err := r.BlockManager.ReadBlock(i)
		if err != nil {
			return fmt.Errorf("Block: Error leyendo bloque %d: %v", i, err)
		}

		// --- Intentar como Directory ---
		var dir file.Directory
		if err := binary.Read(bytes.NewReader(raw), binary.LittleEndian, &dir); err == nil {
			var rows []string
			rows = append(rows, `<TR><TD>b_name</TD><TD>b_inodo</TD></TR>`)

			validEntries := 0
			hasDot, hasDotDot := false, false
			allInodesValid := true

			for _, e := range dir.B_content {
				name := strings.Trim(string(e.B_name[:]), "\x00")
				if e.B_inodo != -1 && name != "" {
					if name == "." {
						hasDot = true
					} else if name == ".." {
						hasDotDot = true
					} else {
						validEntries++
					}
					if e.B_inodo < 0 || e.B_inodo >= sb.S_inodes_count {
						allInodesValid = false
					}
					rows = append(rows, fmt.Sprintf("<TR><TD>%s</TD><TD>%d</TD></TR>", name, e.B_inodo))
				}
			}

			if hasDot && hasDotDot && validEntries > 0 && allInodesValid {
				nodeName := fmt.Sprintf("bloque_carpeta_%d", i)
				nodes = append(nodes, nodeName)

				b.WriteString(fmt.Sprintf("    subgraph cluster_%d {\n", clusterID))
				b.WriteString(fmt.Sprintf("        label = \"Bloque Carpeta %d\";\n", i))
				b.WriteString("        color=white;\n        node [shape=box, margin=0.2, fillcolor=white];\n\n")
				b.WriteString(fmt.Sprintf("        \"%s\" [label=<<TABLE BORDER=\"0\" CELLBORDER=\"0\" CELLSPACING=\"0\">\n", nodeName))
				for _, row := range rows {
					b.WriteString("            " + row + "\n")
				}
				b.WriteString("        </TABLE>>];\n    }\n\n")
				clusterID++
				continue
			}
		}

		// --- Intentar como Pointer ---
		var ptr file.Pointer
		if err := binary.Read(bytes.NewReader(raw), binary.LittleEndian, &ptr); err == nil {
			validPtrs := 0
			var nums []string
			for _, p := range ptr.B_pointers {
				if p != -1 && p < sb.S_blocks_count {
					validPtrs++
				}
				nums = append(nums, fmt.Sprintf("%d", p))
			}
			if validPtrs > 0 {
				nodeName := fmt.Sprintf("bloque_apuntadores_%d", i)
				nodes = append(nodes, nodeName)

				ptrText := wrapText(strings.Join(nums, ", "), 35)
				b.WriteString(fmt.Sprintf("    subgraph cluster_%d {\n", clusterID))
				b.WriteString(fmt.Sprintf("        label = \"Bloque Apuntadores %d\";\n", i))
				b.WriteString("        color=white;\n        node [shape=box, margin=0.2, fillcolor=white];\n\n")
				b.WriteString(fmt.Sprintf("        \"%s\" [label=\"%s\"];\n", nodeName, ptrText))
				b.WriteString("    }\n\n")
				clusterID++
				continue
			}
		}

		// --- Caso por defecto: Contenido ---
		buf := bytes.NewReader(raw)
		var tempArr [64]byte
		if err := binary.Read(buf, binary.LittleEndian, &tempArr); err != nil {
			return fmt.Errorf("Block: Error leyendo bloque %d como raw: %v", i, err)
		}
		content := strings.ReplaceAll(string(tempArr[:]), "\"", "'")
		content = wrapText(content, 35)

		nodeName := fmt.Sprintf("bloque_archivo_%d", i)
		nodes = append(nodes, nodeName)

		b.WriteString(fmt.Sprintf("    subgraph cluster_%d {\n", clusterID))
		b.WriteString(fmt.Sprintf("        label = \"Bloque Archivo %d\";\n", i))
		b.WriteString("        color=white;\n        node [shape=box, margin=0.2, fillcolor=white];\n\n")
		b.WriteString(fmt.Sprintf("        \"%s\" [label=\"%s\"];\n", nodeName, content))
		b.WriteString("    }\n\n")
		clusterID++
	}

	// edges
	for i := 0; i < len(nodes)-1; i++ {
		b.WriteString(fmt.Sprintf("    %s -> %s;\n", nodes[i], nodes[i+1]))
	}

	b.WriteString("}\n")

	const generate_dot bool = false
	err = r.graphvizExecute(b, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("Block: Error generando gráfico: %v", err)
	}

	return nil
}

func (r *ReportSysManager) BitmapInode(path string, format string) error {
	str, err := r.InodeManager.GetBitmapInode(
		r.SBManager.SB.S_bm_inode_start,
		r.SBManager.SB.S_inodes_count,
	)
	if err != nil {
		return err
	}

	// Normalizar: quitar espacios y separar en grupos de 20
	clean := strings.ReplaceAll(str, " ", "")
	var lines []string
	for i := 0; i < len(clean); i += 20 {
		end := i + 20
		if end > len(clean) {
			end = len(clean)
		}
		lines = append(lines, clean[i:end])
	}

	// Calcular resumen
	ocupados := strings.Count(clean, "1")
	libres := strings.Count(clean, "0")

	// Según formato
	switch strings.ToLower(format) {
	case "txt":
		// Escribir archivo TXT
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("BitmapInode: Error creando archivo txt: %v", err)
		}
		defer f.Close()

		for _, line := range lines {
			_, err := f.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("BitmapInode: Error escribiendo en archivo txt: %v", err)
			}
		}
		// Agregar resumen al final
		f.WriteString(fmt.Sprintf("\nResumen: %d ocupados, %d libres\n", ocupados, libres))

	default:
		return fmt.Errorf("BitmapInode: Formato no soportado: %s (solo txt)", format)
	}

	return nil
}

func (r *ReportSysManager) BitmapBlock(path string, format string) error {
	// Obtener string del bitmap
	str, err := r.BlockManager.GetBitmapBlock(r.SBManager.SB.S_bm_block_start, r.SBManager.SB.S_blocks_count)
	if err != nil {
		return err
	}

	// Normalizar: quitar espacios y separar en grupos de 20
	clean := strings.ReplaceAll(str, " ", "")
	var lines []string
	for i := 0; i < len(clean); i += 20 {
		end := i + 20
		if end > len(clean) {
			end = len(clean)
		}
		lines = append(lines, clean[i:end])
	}

	// Calcular resumen
	ocupados := strings.Count(clean, "1")
	libres := strings.Count(clean, "0")

	// Según formato
	switch strings.ToLower(format) {
	case "txt":
		// Escribir archivo TXT
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("BitmapInode: Error creando archivo txt: %v", err)
		}
		defer f.Close()

		for _, line := range lines {
			_, err := f.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("BitmapInode: Error escribiendo en archivo txt: %v", err)
			}
		}

		f.WriteString(fmt.Sprintf("\nResumen: %d ocupados, %d libres\n", ocupados, libres))

	default:
		return fmt.Errorf("BitmapInode: Formato no soportado: %s (solo txt)", format)
	}

	return nil
}

func (r *ReportSysManager) Tree(path string, format string) error {
	sb, err := r.SBManager.ReadSuperBlock()
	if err != nil {
		return fmt.Errorf("Tree: Error leyendo superblock: %v", err)
	}

	rootID := sb.S_first_ino

	var builder strings.Builder
	builder.WriteString("digraph G {\n")
	builder.WriteString("  rankdir=LR;\n")
	builder.WriteString("  node [fontname=\"Helvetica\"];\n")

	visited := make(map[int64]bool)

	var walk func(inodeID int64, parentName string) error
	walk = func(inodeID int64, parentName string) error {
		if visited[inodeID] {
			return nil
		}
		visited[inodeID] = true

		inode, err := r.InodeManager.ReadInode(inodeID)
		if err != nil {
			return err
		}

		// ----- Inodo -----
		inodeNode := fmt.Sprintf("inode_%d", inodeID)
		tipo := map[byte]string{0: "FILE", 1: "DIR"}[inode.I_type]

		builder.WriteString(fmt.Sprintf("subgraph cluster_inode_%d {\n", inodeID))
		builder.WriteString(fmt.Sprintf("  label = \"Inodo %d\";\n", inodeID))
		builder.WriteString("  color=white;\n")
		builder.WriteString("  node [shape=box, margin=0.2, fillcolor=white];\n")
		builder.WriteString(fmt.Sprintf("  \"%s\" [fillcolor=\"#EAD8A4\", color=white, style=filled, label=<<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"5\" BGCOLOR=\"#EAD8A4\">", inodeNode))
		builder.WriteString(fmt.Sprintf("<TR><TD>i_uid</TD><TD>%d</TD></TR>", inode.I_uid))
		builder.WriteString(fmt.Sprintf("<TR><TD>i_gid</TD><TD>%d</TD></TR>", inode.I_gid))
		builder.WriteString(fmt.Sprintf("<TR><TD>size</TD><TD>%d</TD></TR>", inode.I_size))
		builder.WriteString(fmt.Sprintf("<TR><TD>i_type</TD><TD>%s</TD></TR>", tipo))
		builder.WriteString(fmt.Sprintf("<TR><TD>i_permissions</TD><TD>%d</TD></TR>", inode.I_permissions))
		// Array de bloques
		var bloques []string
		for _, blk := range inode.I_block {
			bloques = append(bloques, fmt.Sprintf("%d", blk))
		}
		builder.WriteString(fmt.Sprintf("<TR><TD>Bloques</TD><TD>[%s]</TD></TR>", strings.Join(bloques, ", ")))
		builder.WriteString("</TABLE>>];\n")
		builder.WriteString("}\n")

		// Conectar con el padre
		if parentName != "" {
			builder.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", parentName, inodeNode))
		}

		// ----- Recorrer bloques -----
		for i, blockID := range inode.I_block {
			if blockID == -1 {
				continue
			}

			if i < 12 { // Bloques directos
				if inode.I_type == DIR_TYPE {
					// Bloque Carpeta
					blockNode := fmt.Sprintf("bloque_carpeta_%d", blockID)
					entries, err := r.DirectoryManager.ReadDirectory(inodeID)
					if err != nil {
						return err
					}

					builder.WriteString(fmt.Sprintf("subgraph cluster_block_%d {\n", blockID))
					builder.WriteString(fmt.Sprintf("  label = \"Bloque Carpeta %d\";\n", blockID))
					builder.WriteString("  color=white;\n")
					builder.WriteString("  node [shape=box, margin=0.2, fillcolor=white];\n")
					builder.WriteString(fmt.Sprintf("  \"%s\" [fillcolor=\"#ECEDB0\", color=white, style=filled, label=<<TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"5\" BGCOLOR=\"#ECEDB0\">", blockNode))
					builder.WriteString("<TR><TD>b_name</TD><TD>b_inodo</TD></TR>")
					for _, e := range entries {
						name := strings.Trim(string(e.B_name[:]), "\x00")
						if name == "" || name == "." || name == ".." {
							continue
						}
						builder.WriteString(fmt.Sprintf("<TR><TD>%s</TD><TD>%d</TD></TR>", name, e.B_inodo))
					}
					builder.WriteString("</TABLE>>];\n")
					builder.WriteString("}\n")

					builder.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", inodeNode, blockNode))

					// Recursión a hijos
					for _, e := range entries {
						name := strings.Trim(string(e.B_name[:]), "\x00")
						if name == "" || name == "." || name == ".." {
							continue
						}
						if err := walk(e.B_inodo, blockNode); err != nil {
							return err
						}
					}

				} else {

					// Bloque Archivo
					blockNode := fmt.Sprintf("bloque_archivo_%d", blockID)
					data, _ := r.FileManager.ReadFile(blockID)
					content := wrapText(strings.ReplaceAll(string(data), "\n", "\\n"), 60)

					builder.WriteString(fmt.Sprintf("subgraph cluster_block_%d {\n", blockID))
					builder.WriteString(fmt.Sprintf("  label = \"Bloque Archivo %d\";\n", blockID))
					builder.WriteString("  color=white;\n")
					builder.WriteString("  node [shape=box, margin=0.2, fillcolor=white];\n")
					builder.WriteString(fmt.Sprintf("  \"%s\" [fillcolor=\"#EAD8A4\", color=white, style=filled, label=\"%s\"];\n", blockNode, content))
					builder.WriteString("}\n")

					builder.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", inodeNode, blockNode))
				}

			} else if i == 12 { // Indirecto simple
				indNode := fmt.Sprintf("bloque_apuntadores_%d", blockID)
				ptrs, err := r.BlockManager.ReadPointers(blockID)
				if err != nil {
					return err
				}

				builder.WriteString(fmt.Sprintf("subgraph cluster_block_%d {\n", blockID))
				builder.WriteString(fmt.Sprintf("  label = \"Bloque Apuntadores %d\";\n", blockID))
				builder.WriteString("  color=white;\n")
				builder.WriteString("  node [shape=box, margin=0.2, fillcolor=white];\n")
				builder.WriteString(fmt.Sprintf("  \"%s\" [fillcolor=\"#F68537\", style=filled, fontcolor=white, label=\"%v\"];\n", indNode, ptrs))
				builder.WriteString("}\n")

				builder.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", inodeNode, indNode))
			}
		}

		return nil
	}

	// Nodo raíz
	rootName := fmt.Sprintf("root_%d", rootID)
	builder.WriteString(fmt.Sprintf("  \"%s\" [label=\"/\", shape=folder, style=filled, fillcolor=lightgreen];\n", rootName))

	if err := walk(rootID, rootName); err != nil {
		return err
	}

	builder.WriteString("}\n")

	const generate_dot bool = false
	err = r.graphvizExecute(builder, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("Tree: Error generando gráfico: %v", err)
	}

	return nil
}

func (r *ReportSysManager) SuperBlock(path string, format string) error {
	sb, err := r.SBManager.ReadSuperBlock()
	if err != nil {
		return fmt.Errorf("ReporteSuperBlock: Error leyendo SuperBlock %w", err)
	}

	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("    graph [rankdir=TB, bgcolor=\"white\"];\n")
	b.WriteString("    node [shape=plain, fontname=\"Times New Roman\", fontsize=10];\n")
	b.WriteString("    edge [style=invis];\n\n")
	b.WriteString("    reporte_sb [label=<\n")
	b.WriteString("        <TABLE BORDER=\"1\" CELLBORDER=\"0.1\" CELLSPACING=\"1\" CELLPADDING=\"2\" BGCOLOR=\"white\">\n")
	b.WriteString("            <TR><TD COLSPAN=\"2\" BGCOLOR=\"#2E8B57\" ALIGN=\"Center\">\n")
	b.WriteString("                <FONT COLOR=\"white\" POINT-SIZE=\"12\"><B>REPORTE SUPERBLOCK</B></FONT>\n")
	b.WriteString("            </TD></TR>\n")

	// --- Campos del SuperBlock ---
	b.WriteString(tdRow("s_filesystem_type", cleanCString(sb.S_filesystem_type[:]), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_inodes_count", fmt.Sprintf("%d", sb.S_inodes_count), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_blocks_count", fmt.Sprintf("%d", sb.S_blocks_count), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_free_blocks_count", fmt.Sprintf("%d", sb.S_free_blocks_count), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_free_inodes_count", fmt.Sprintf("%d", sb.S_free_inodes_count), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_mtime", formatTimeFromBytes(sb.S_mtime[:]), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_umtime", formatTimeFromBytes(sb.S_umtime[:]), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_mnt_count", fmt.Sprintf("%d", sb.S_mnt_count), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_magic", fmt.Sprintf("%d", sb.S_magic), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_inode_size", fmt.Sprintf("%d", sb.S_inode_size), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_block_size", fmt.Sprintf("%d", sb.S_block_size), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_first_ino", fmt.Sprintf("%d", sb.S_first_ino), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_first_blo", fmt.Sprintf("%d", sb.S_first_blo), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_bm_inode_start", fmt.Sprintf("%d", sb.S_bm_inode_start), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_bm_block_start", fmt.Sprintf("%d", sb.S_bm_block_start), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_inode_start", fmt.Sprintf("%d", sb.S_inode_start), "#C9E6D3", "#E6F2EC"))
	b.WriteString(tdRow("s_block_start", fmt.Sprintf("%d", sb.S_block_start), "#C9E6D3", "#E6F2EC"))

	b.WriteString("            <TR><TD COLSPAN=\"2\" ALIGN=\"CENTER\"><FONT POINT-SIZE=\"8\">Reporte de SuperBlock</FONT></TD></TR>\n")
	b.WriteString("        </TABLE>\n")
	b.WriteString("    >];\n")
	b.WriteString("}\n")

	const generate_dot bool = false

	err = r.graphvizExecute(b, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("Tree: Error generando gráfico: %v", err)
	}

	return nil
}

func (r *ReportSysManager) File(userPath string, virtualPath string, format string) error {
	// Validar formato
	if format != "txt" {
		return fmt.Errorf("File: Formato '%s' no soportado, solo se permite 'txt'", format)
	}

	// Normalizar ruta del archivo en FS virtual
	virtualPath = path.Clean("/" + strings.Trim(virtualPath, "/"))

	// Resolver directorio padre
	parentID, name, err := r.DirectoryManager.ResolvePath(virtualPath)
	if err != nil {
		return fmt.Errorf("File: No se pudo resolver ruta '%s': %v", virtualPath, err)
	}

	// Buscar inodo del archivo
	fileInodeID, err := r.DirectoryManager.FindEntry(parentID, name)
	if err != nil {
		return fmt.Errorf("File: Archivo '%s' no encontrado", virtualPath)
	}

	// Leer inodo
	inode, err := r.InodeManager.ReadInode(fileInodeID)
	if err != nil {
		return fmt.Errorf("File: Error leyendo inodo %d: %v", fileInodeID, err)
	}
	if inode.I_type != FILE_TYPE { // asumiendo FILE_TYPE = 1
		return fmt.Errorf("File: '%s' no es un archivo", virtualPath)
	}

	// Leer contenido
	content, err := r.FileManager.ReadFile(fileInodeID)
	if err != nil {
		return fmt.Errorf("File: Error leyendo archivo '%s': %v", virtualPath, err)
	}

	// Armar salida en texto plano
	var b strings.Builder
	b.WriteString("Reporte de archivo\n")
	b.WriteString(fmt.Sprintf("Archivo: %s\n", name))
	b.WriteString("Contenido:\n")
	b.WriteString(string(content))

	// Crear carpeta destino si no existe
	if err := os.MkdirAll(filepath.Dir(userPath), 0o755); err != nil {
		return fmt.Errorf("File: No se pudo crear directorio destino: %v", err)
	}

	// Escribir archivo .txt
	if err := os.WriteFile(userPath, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("File: Error escribiendo reporte '%s': %v", userPath, err)
	}

	return nil
}

func (r *ReportSysManager) Ls(userPath string, virtualPath string, format string) error {
	// Normalizar ruta
	virtualPath = path.Clean("/" + strings.Trim(virtualPath, "/"))

	// Resolver directorio
	parentID, name, err := r.DirectoryManager.ResolvePath(virtualPath)
	if err != nil {
		return fmt.Errorf("Ls: No se pudo resolver ruta '%s': %v", virtualPath, err)
	}

	// Determinar inode del directorio
	var dirInodeID int64
	if virtualPath == "/" {
		dirInodeID = parentID
	} else {
		dirInodeID, err = r.DirectoryManager.FindEntry(parentID, name)
		if err != nil {
			return fmt.Errorf("Ls: Directorio '%s' no encontrado", virtualPath)
		}
	}

	// Leer inodo
	inode, err := r.InodeManager.ReadInode(dirInodeID)
	if err != nil {
		return fmt.Errorf("Ls: Error leyendo inodo %d: %v", dirInodeID, err)
	}
	if inode.I_type != DIR_TYPE {
		return fmt.Errorf("Ls: '%s' no es un directorio", virtualPath)
	}

	// Listar entradas
	entries, err := r.DirectoryManager.ReadDirectory(dirInodeID)
	if err != nil {
		return fmt.Errorf("Ls: Error leyendo directorio: %v", err)
	}

	// --- Construcción del DOT ---
	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("  node [shape=plaintext fontname=\"Courier\"];\n")
	b.WriteString(fmt.Sprintf("  \"%s\" [label=<\n", virtualPath))
	b.WriteString("    <TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n")
	b.WriteString("      <TR><TD COLSPAN=\"8\"><B>LS Report</B></TD></TR>\n")
	b.WriteString(fmt.Sprintf("      <TR><TD COLSPAN=\"8\">Directorio: %s</TD></TR>\n", virtualPath))
	b.WriteString("      <TR>")
	b.WriteString("<TD><B>Permisos</B></TD><TD><B>Owner</B></TD><TD><B>Grupo</B></TD>")
	b.WriteString("<TD><B>Size (en Bytes)</B></TD><TD><B>Fecha</B></TD><TD><B>Hora</B></TD>")
	b.WriteString("<TD><B>Tipo</B></TD><TD><B>Name</B></TD>")
	b.WriteString("</TR>\n")

	// Recorrer entradas y detallar info
	for _, entry := range entries {
		name := strings.Trim(string(entry.B_name[:]), "\x00")
		if name == "" {
			continue
		}

		// Leer inodo de la entrada
		childInode, err := r.InodeManager.ReadInode(entry.B_inodo)
		if err != nil {
			return fmt.Errorf("Ls: Error leyendo inodo hijo %d: %v", entry.B_inodo, err)
		}

		// Permisos en estilo rwx
		perms := r.permsToString(childInode.I_permissions, childInode.I_type)

		// Dueño y grupo (por ahora solo números, a menos que quieras mapear desde /users.txt)
		owner := fmt.Sprintf("UID:%d", childInode.I_uid)
		group := fmt.Sprintf("GID:%d", childInode.I_gid)

		// Tamaño
		size := childInode.I_size

		// Fecha y hora (suponiendo que tienes timestamps en el inodo, si no, dejamos vacío)
		fecha := time.Unix(childInode.I_ctime, 0).Format("02/01/2006")
		hora := time.Unix(childInode.I_ctime, 0).Format("15:04")

		// Tipo
		tipo := "Archivo"
		if childInode.I_type == DIR_TYPE {
			tipo = "Carpeta"
		}

		// Agregar fila
		b.WriteString(fmt.Sprintf(
			"      <TR><TD>%s</TD><TD>%s</TD><TD>%s</TD><TD>%d</TD><TD>%s</TD><TD>%s</TD><TD>%s</TD><TD>%s</TD></TR>\n",
			perms, owner, group, size, fecha, hora, tipo, name,
		))
	}

	b.WriteString("    </TABLE>\n")
	b.WriteString("  >];\n")
	b.WriteString("}\n")

	const generate_dot bool = false

	// Generar salida con Graphviz
	if err := r.graphvizExecute(b, userPath, format, generate_dot); err != nil {
		return fmt.Errorf("Ls: Error generando reporte con Graphviz: %v", err)
	}

	return nil
}

/*

HELPERS

*/

func (r *ReportSysManager) permsToString(perms uint16, fileType byte) string {
	var sb strings.Builder

	// Tipo
	if fileType == DIR_TYPE {
		sb.WriteByte('d')
	} else {
		sb.WriteByte('-')
	}

	// Dueño
	if perms&0o400 != 0 {
		sb.WriteByte('r')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o200 != 0 {
		sb.WriteByte('w')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o100 != 0 {
		sb.WriteByte('x')
	} else {
		sb.WriteByte('-')
	}

	// Grupo
	if perms&0o040 != 0 {
		sb.WriteByte('r')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o020 != 0 {
		sb.WriteByte('w')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o010 != 0 {
		sb.WriteByte('x')
	} else {
		sb.WriteByte('-')
	}

	// Otros
	if perms&0o004 != 0 {
		sb.WriteByte('r')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o002 != 0 {
		sb.WriteByte('w')
	} else {
		sb.WriteByte('-')
	}
	if perms&0o001 != 0 {
		sb.WriteByte('x')
	} else {
		sb.WriteByte('-')
	}

	return sb.String()
}

func (r *ReportSysManager) graphvizExecute(b strings.Builder, userPath string, format string, generateDot bool) error {
	// userPath: viene con extensión final (.jpg, .png, etc.)
	outputImgPath := userPath
	dotPath := strings.TrimSuffix(userPath, filepath.Ext(userPath)) + ".dot"

	// 1) Crear carpeta destino
	if err := os.MkdirAll(filepath.Dir(dotPath), 0o755); err != nil {
		return fmt.Errorf("graphvizExecute: No se pudo crear directorio destino: %w", err)
	}

	// 2) Escribir archivo DOT
	if err := os.WriteFile(dotPath, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("graphvizExecute: No se pudo escribir DOT: %w", err)
	}

	// 3) Si el usuario pidió DOT explícitamente
	if format == "dot" {
		return nil
	}

	// 4) Validar formatos soportados
	supported := map[string]bool{
		"png": true,
		"pdf": true,
		"svg": true,
		"jpg": true,
	}
	if !supported[format] {
		return fmt.Errorf("graphvizExecute: Formato %q no soportado. Se generó el archivo DOT en: %s", format, dotPath)
	}

	// 5) Asegurar directorios de salida
	if err := os.MkdirAll(filepath.Dir(outputImgPath), 0o755); err != nil {
		return fmt.Errorf("graphvizExecute: Error creando directorios de salida: %v", err)
	}

	// 6) Ejecutar Graphviz
	cmd := exec.Command("dot", "-T"+format, dotPath, "-o", outputImgPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"graphvizExecute: Error ejecutando Graphviz: %v\nDetalles: %s\n(DOT generado en %s)",
			err, stderr.String(), dotPath,
		)
	}

	// 7) Si no quieres conservar el archivo DOT
	if !generateDot {
		err := os.Remove(dotPath)
		if err != nil {
			return fmt.Errorf("graphvizExecute: Error eliminando archivo DOT: %v", err)
		}
	}

	return nil
}

func wrapText(s string, n int) string {
	var out strings.Builder
	for i, r := range s {
		out.WriteRune(r)
		if (i+1)%n == 0 {
			out.WriteRune('\n')
		}
	}
	return out.String()
}

func tdRow(label, value, bgLabel, bgValue string) string {
	return fmt.Sprintf(
		"            <TR>\n"+
			"                <TD BGCOLOR=\"%s\" ALIGN=\"Center\">%s</TD>\n"+
			"                <TD BGCOLOR=\"%s\" ALIGN=\"left\">%s</TD>\n"+
			"            </TR>\n", bgLabel, label, bgValue, value)
}

func cleanCString(b []byte) string {
	return strings.TrimRight(strings.Trim(string(b), "\x00"), " ")
}

func formatTimeFromBytes(b []byte) string {
	// Si guardaste texto tipo "YYYY-MM-DD HH:MM", solo limpia nulos.
	s := cleanCString(b)
	if s == "" {
		return "-"
	}
	return s
}
