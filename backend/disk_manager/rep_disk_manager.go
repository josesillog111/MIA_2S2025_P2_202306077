package dmanager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ReportDiskManager struct {
	Partition   MountedPartition
	DiskManager *DiskManager
}

func NewReportDiskManager(partition MountedPartition) *ReportDiskManager {
	return &ReportDiskManager{
		Partition:   partition,
		DiskManager: NewDiskManager(),
	}
}
func calcularPorcentaje(size, total int64) string {
	if total == 0 {
		return "0%"
	}
	percent := float64(size) / float64(total) * 100
	return fmt.Sprintf("%.0f%%", percent)
}

func graphvizExecute(b strings.Builder, userPath string, format string, generateDot bool) error {
	// userPath: viene con extensión final (.jpg, .png, etc.)
	outputImgPath := userPath
	dotPath := strings.TrimSuffix(userPath, filepath.Ext(userPath)) + ".dot"

	// 1) Crear carpeta destino
	if err := os.MkdirAll(filepath.Dir(dotPath), 0o755); err != nil {
		return fmt.Errorf("no se pudo crear directorio destino: %w", err)
	}

	// 2) Escribir archivo DOT
	if err := os.WriteFile(dotPath, []byte(b.String()), 0o644); err != nil {
		return fmt.Errorf("no se pudo escribir DOT: %w", err)
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
		return fmt.Errorf("GraphvizExecute: Formato %q no soportado. Se generó el archivo DOT en: %s", format, dotPath)
	}

	// 5) Asegurar directorios de salida
	if err := os.MkdirAll(filepath.Dir(outputImgPath), 0o755); err != nil {
		return fmt.Errorf("GraphvizExecute: Error creando directorios de salida: %v", err)
	}

	// 6) Ejecutar Graphviz
	cmd := exec.Command("dot", "-T"+format, dotPath, "-o", outputImgPath)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"GraphvizExecute: Error ejecutando Graphviz: %v\nDetalles: %s\n(DOT generado en %s)",
			err, stderr.String(), dotPath,
		)
	}

	// 7) Si no quieres conservar el archivo DOT
	if !generateDot {
		err := os.Remove(dotPath)
		if err != nil {
			return fmt.Errorf("GraphvizExecute: Error eliminando archivo DOT: %v", err)
		}
	}

	return nil
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

func colorHeader(bg string, title string) string {
	return fmt.Sprintf(
		"            <TR>\n"+
			"                <TD COLSPAN=\"2\" BGCOLOR=\"%s\" ALIGN=\"Center\">\n"+
			"                    <FONT COLOR=\"white\" FACE=\"Times New Roman\" POINT-SIZE=\"10\"><B>%s</B></FONT>\n"+
			"                </TD>\n"+
			"            </TR>\n",
		bg, title,
	)
}

func (r *ReportDiskManager) Mbr(path string, id string, format string) error {
	mbr, err := r.DiskManager.MbrManager.ReadMBR(r.Partition.Path)
	if err != nil {
		return fmt.Errorf("ReporteMBR: Error leyendo MBR %w", err)
	}
	var b strings.Builder
	b.WriteString("digraph G {\n")
	b.WriteString("    graph [rankdir=TB, bgcolor=\"white\"];\n")
	b.WriteString("    node [shape=plain, fontname=\"Times New Roman\", fontsize=10];\n")
	b.WriteString("    edge [style=invis];\n\n")
	b.WriteString("    reporte_mbr [label=<\n")
	b.WriteString("        <TABLE BORDER=\"1\" CELLBORDER=\"0.1\" CELLSPACING=\"1\" CELLPADDING=\"2\" BGCOLOR=\"white\">\n")
	b.WriteString("            <TR><TD COLSPAN=\"2\" BGCOLOR=\"#6A327F\" ALIGN=\"Center\">\n")
	b.WriteString("                <FONT COLOR=\"white\" POINT-SIZE=\"12\"><B>REPORTE DE MBR</B></FONT>\n")
	b.WriteString("            </TD></TR>\n")

	b.WriteString(tdRow("mbr_tamano", fmt.Sprintf("%d", mbr.Mbr_tamano), "#E6DAF0", "#D8C2E6"))
	b.WriteString(tdRow("mbr_fecha_creacion", formatTimeFromBytes(mbr.Mbr_fecha_creacion[:]), "#E6DAF0", "#D8C2E6"))
	b.WriteString(tdRow("mbr_disk_signature", fmt.Sprintf("%d", mbr.Mbr_disk_signature), "#E6DAF0", "#D8C2E6"))

	for i := 0; i < 4; i++ {
		part := mbr.Mbr_partition[i]
		if part.Part_size <= 0 {
			continue
		}
		b.WriteString(colorHeader("#B07CBE", "Particion"))
		b.WriteString(tdRow("part_status", fmt.Sprintf("%d", part.Part_status), "#E6DAF0", "#D8C2E6"))
		b.WriteString(tdRow("part_type", strings.ToLower(string(part.Part_type)), "#E6DAF0", "#D8C2E6"))
		b.WriteString(tdRow("part_fit", strings.ToLower(string(part.Part_fit)), "#E6DAF0", "#D8C2E6"))
		b.WriteString(tdRow("part_start", fmt.Sprintf("%d", part.Part_start), "#E6DAF0", "#D8C2E6"))
		b.WriteString(tdRow("part_size", fmt.Sprintf("%d", part.Part_size), "#E6DAF0", "#D8C2E6"))
		b.WriteString(tdRow("part_name", cleanCString(part.Part_name[:]), "#E6DAF0", "#D8C2E6"))

		if part.Part_type == 'E' || part.Part_type == 'e' {
			ebrs, err := r.DiskManager.EbrManager.GetEBRs(r.Partition.Path, part)
			if err != nil {
				return fmt.Errorf("ReporteMBR: Error leyendo EBRs %w", err)
			}
			for _, ebr := range ebrs {
				b.WriteString(colorHeader("#F28E8C", "Particion Logica"))
				b.WriteString(tdRow("part_status", fmt.Sprintf("%d", ebr.Part_mount), "#F7D0CF", "#F5B9B7"))
				b.WriteString(tdRow("part_next", fmt.Sprintf("%d", ebr.Part_next), "#F7D0CF", "#F5B9B7"))
				b.WriteString(tdRow("part_fit", strings.ToLower(string(ebr.Part_fit)), "#F7D0CF", "#F5B9B7"))
				b.WriteString(tdRow("part_start", fmt.Sprintf("%d", ebr.Part_start), "#F7D0CF", "#F5B9B7"))
				b.WriteString(tdRow("part_size", fmt.Sprintf("%d", ebr.Part_size), "#F7D0CF", "#F5B9B7"))
				b.WriteString(tdRow("part_name", cleanCString(ebr.Part_name[:]), "#F7D0CF", "#F5B9B7"))
			}
		}
	}

	b.WriteString("            <TR><TD COLSPAN=\"2\" ALIGN=\"CENTER\"><FONT POINT-SIZE=\"8\">Reporte de MBR</FONT></TD></TR>\n")
	b.WriteString("        </TABLE>\n")
	b.WriteString("    >];\n")
	b.WriteString("}\n")

	const generate_dot bool = false

	// 4) Ejecutar Graphviz
	err = graphvizExecute(b, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("ReporteMBR: Error generando imagen: %w", err)
	}

	return nil
}

func (r *ReportDiskManager) Disk(path string, id string, format string) error {
	mbr, err := r.DiskManager.MbrManager.ReadMBR(r.Partition.Path)
	if err != nil {
		return fmt.Errorf("Disk: Error leyendo MBR %w", err)
	}

	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	sb.WriteString("    node [shape=none, fontname=\"Helvetica\", fontsize=10];\n\n")
	sb.WriteString("    disco [label=<\n")
	sb.WriteString("        <TABLE BORDER=\"1\" CELLBORDER=\"0.5\" CELLSPACING=\"3\" CELLPADDING=\"0\" COLOR=\"lightblue\">\n")
	sb.WriteString(fmt.Sprintf("            <TR><TD BORDER=\"0\" ALIGN=\"CENTER\" COLSPAN=\"5\"><B>%s</B></TD></TR>\n", filepath.Base(r.Partition.Path)))
	sb.WriteString("            <TR>\n")

	// MBR fijo
	sb.WriteString("                <TD ALIGN=\"CENTER\" BORDER=\"1\" COLOR=\"lightblue\" HEIGHT=\"50\" WIDTH=\"80\"><B>MBR</B></TD>\n")

	var usedDisk int64
	particiones := mbr.Mbr_partition

	// Iteramos particiones
	for i := 0; i < 4; i++ {
		part := particiones[i]
		if part.Part_size <= 0 {
			continue
		}

		usedDisk += part.Part_size
		porcentaje := calcularPorcentaje(part.Part_size, mbr.Mbr_tamano)

		switch part.Part_type {
		case 'P', 'p':
			// Primaria
			sb.WriteString(fmt.Sprintf("                <TD ALIGN=\"CENTER\" BORDER=\"1\" COLOR=\"lightblue\" HEIGHT=\"50\" WIDTH=\"100\"><B>Primaria</B><BR/>%s del disco</TD>\n", porcentaje))

		case 'E', 'e':
			// Extendida con tabla interna
			sb.WriteString("                <TD BORDER=\"1\" COLOR=\"lightblue\" ALIGN=\"CENTER\" VALIGN=\"TOP\" HEIGHT=\"50\" WIDTH=\"300\">\n")
			sb.WriteString("                    <TABLE BORDER=\"0\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"4\" COLOR=\"lightblue\">\n")

			ebrList, err := r.DiskManager.EbrManager.GetEBRs(r.Partition.Path, part)
			if err != nil {
				return fmt.Errorf("Disk: Error leyendo EBRs %w", err)
			}

			// Adjust the colspan for the extended partition header
			colspan := 1
			if len(ebrList) > 0 {
				colspan = 2 * len(ebrList)
			}
			sb.WriteString(fmt.Sprintf("                        <TR><TD ALIGN=\"CENTER\" COLSPAN=\"%d\"><B>Extendida</B><BR/>%s del disco</TD></TR>\n", colspan, porcentaje))

			// Now, we create a single row for all the logical and ebr partitions
			sb.WriteString("                        <TR>\n")

			if len(ebrList) == 0 {
				sb.WriteString("                            <TD ALIGN=\"CENTER\" HEIGHT=\"50\"><B>Lógica Libre</B><BR/>100% de la extendida</TD>\n")
			} else {
				for _, ebr := range ebrList {
					porcL := calcularPorcentaje(ebr.Part_size, mbr.Mbr_tamano)
					sb.WriteString("                            <TD ALIGN=\"CENTER\" HEIGHT=\"50\"><B>EBR</B></TD>\n")
					sb.WriteString(fmt.Sprintf("                            <TD ALIGN=\"CENTER\" HEIGHT=\"50\"><B>Lógica</B><BR/>%s</TD>\n", porcL))
				}
			}
			sb.WriteString("                        </TR>\n")

			sb.WriteString("                    </TABLE>\n")
			sb.WriteString("                </TD>\n")
		}
	}

	// Espacio libre al final
	restante := mbr.Mbr_tamano - usedDisk
	if restante < 0 {
		restante = 0
	}
	porcRestante := calcularPorcentaje(restante, mbr.Mbr_tamano)
	if restante > 0 {
		sb.WriteString(fmt.Sprintf("                <TD ALIGN=\"CENTER\" BORDER=\"1\" COLOR=\"lightblue\" HEIGHT=\"50\" WIDTH=\"80\"><B>Libre</B><BR/>%s del disco</TD>\n", porcRestante))
	}

	sb.WriteString("            </TR>\n")
	sb.WriteString("        </TABLE>\n")
	sb.WriteString("    >];\n")
	sb.WriteString("}\n")

	const generate_dot bool = false

	// 4) Ejecutar Graphviz
	err = graphvizExecute(sb, path, format, generate_dot)
	if err != nil {
		return fmt.Errorf("Disk: Error generando imagen: %w", err)
	}
	return nil
}
