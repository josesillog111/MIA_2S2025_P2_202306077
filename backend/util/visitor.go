package util

import (
	dmanager "backend/disk_manager"
	ext2 "backend/ext2_manager"
	ext3 "backend/ext3_manager"
	fsys "backend/file_system"
	parser "backend/parser"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// fileRegex matches parameters like -file1, -file2, etc.
var fileRegex = regexp.MustCompile(`-file(\d+)`)

func (v *Visitor) NewFileSystem(fsType string, partitionId string) (fsys.FileSystem, error) {
	switch strings.ToLower(fsType) {
	case "2fs":
		partition := v.ListMountedPartitions.GetPartitionById(partitionId)
		if partition == nil {
			return nil, fmt.Errorf("NewFileSystem: la partición con id '%s' no está montada", partitionId)
		}
		return ext2.NewEXT2(*partition), nil
	case "3fs":
		partition := v.ListMountedPartitions.GetPartitionById(partitionId)
		if partition == nil {
			return nil, fmt.Errorf("NewFileSystem: la partición con id '%s' no está montada", partitionId)
		}
		return ext3.NewEXT3(*partition), nil
	default:
		return nil, fmt.Errorf("NewFileSystem: Sistema de archivos '%s' no soportado", fsType)
	}
}

func extractValue(strLit antlr.TerminalNode, unquoted antlr.TerminalNode) string {
	if strLit != nil {
		text := strLit.GetText()
		// quitar comillas
		if len(text) > 1 && (text[0] == '"' || text[0] == '\'') {
			return text[1 : len(text)-1]
		}
		return text
	}
	if unquoted != nil {
		return unquoted.GetText()
	}
	return ""
}

func cleanQuotes(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') ||
			(s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// extractInt tries to extract an integer from an antlr.Tree node.
func extractInt(node antlr.Tree) int {
	if t, ok := node.(antlr.TerminalNode); ok {
		text := t.GetText()
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
	}
	return 0
}

type Visitor struct {
	Console               string
	Errors                string
	DiskManager           *dmanager.DiskManager
	ListMountedPartitions *dmanager.MountedPartitionList
	IdMountedAndLogued    string /* ID de la partición montada y del usuario logueado */
	parser.BaseGoDiskGrammarVisitor
}

func NewVisitor() *Visitor {
	return &Visitor{
		Console:               "",
		Errors:                "",
		DiskManager:           dmanager.NewDiskManager(),
		ListMountedPartitions: dmanager.NewMountedPartitionList(),
		IdMountedAndLogued:    "",
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}

	tree.Accept(v)
	return nil
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) interface{} {
	if v.Errors != "" { // detener la ejecución del resto de instrucciones en caso de error
		return nil
	}
	for _, stmt := range ctx.AllStmt() {
		switch n := stmt.(type) {
		case *parser.MKDISKContext:
			v.VisitMKDISK(n)
		case *parser.RMDISKContext:
			v.VisitRMDISK(n)
		case *parser.FDISKContext:
			v.VisitFDISK(n)
		case *parser.MOUNTContext:
			v.VisitMOUNT(n)
		case *parser.MOUNTEDContext:
			v.VisitMOUNTED(n)
		case *parser.MKFSContext:
			v.VisitMKFS(n)
		case *parser.CATContext:
			v.VisitCAT(n)
		case *parser.LOGINContext:
			v.VisitLOGIN(n)
		case *parser.LOGOUTContext:
			v.VisitLOGOUT(n)
		case *parser.MKGRPContext:
			v.VisitMKGRP(n)
		case *parser.RMGRPContext:
			v.VisitRMGRP(n)
		case *parser.MKUSRContext:
			v.VisitMKUSR(n)
		case *parser.RMUSRContext:
			v.VisitRMUSR(n)
		case *parser.CHGRPContext:
			v.VisitCHGRP(n)
		case *parser.MKFILEContext:
			v.VisitMKFILE(n)
		case *parser.MKDIRContext:
			v.VisitMKDIR(n)
		case *parser.REMOVEContext:
			v.VisitREMOVE(n)
		case *parser.EDITContext:
			v.VisitEDIT(n)
		case *parser.RENAMEContext:
			v.VisitRENAME(n)
		case *parser.COPYContext:
			v.VisitCOPY(n)
		case *parser.MOVEContext:
			v.VisitMOVE(n)
		case *parser.FINDContext:
			v.VisitFIND(n)
		case *parser.CHOWNContext:
			v.VisitCHOWN(n)
		case *parser.CHMODContext:
			v.VisitCHMOD(n)
		case *parser.RECOVERYContext:
			v.VisitRECOVERY(n)
		case *parser.LOSSContext:
			v.VisitLOSS(n)
		case *parser.JOURNALINGContext:
			v.VisitJOURNALING(n)
		case *parser.REPContext:
			v.VisitREP(n)
		default:
			v.Errors += fmt.Sprintf("Error: Comando %s desconocido", stmt.GetText())
			return nil
		}
		if v.Errors != "" {
			break
		}
	}
	return nil
}

// Crear un nuevo disco
func (v *Visitor) VisitMKDISK(ctx *parser.MKDISKContext) interface{} {
	// Valores por defecto
	var size int64
	var path string
	fit := "FF"
	unit := "M"

	// Para controlar duplicados
	used := make(map[string]bool)

	// Recorremos todos los parámetros
	for _, param := range ctx.Mkdisk_params().AllMkdisk_param() {
		switch {
		case param.Size() != nil:
			if used["size"] {
				v.Errors += "MKDISK: El parámetro 'size' está duplicado.\n"
				return nil
			}
			used["size"] = true

			textSize := param.Size().INT_LIT().GetText()
			parsedSize, err := strconv.ParseInt(textSize, 10, 64)
			if err != nil || parsedSize <= 0 {
				v.Errors += "MKDISK: Tamaño debe ser positivo.\n"
				return nil
			}
			size = parsedSize

		case param.Path() != nil:
			if used["path"] {
				v.Errors += "MKDISK: El parámetro 'path' está duplicado.\n"
				return nil
			}
			used["path"] = true

			path = extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT())
			if !strings.HasSuffix(strings.ToLower(path), ".mia") {
				path += ".mia"
			}

		case param.Fit() != nil:
			if used["fit"] {
				v.Errors += "MKDISK: El parámetro 'fit' está duplicado.\n"
				return nil
			}
			used["fit"] = true

			fit = strings.ToUpper(param.Fit().ID().GetText())
			if fit != "BF" && fit != "FF" && fit != "WF" {
				v.Errors += "MKDISK: Tipo de ajuste inválido. Use BF, FF o WF.\n"
				return nil
			}

		case param.Unit() != nil:
			if used["unit"] {
				v.Errors += "MKDISK: El parámetro 'unit' está duplicado.\n"
				return nil
			}
			used["unit"] = true

			unit = strings.ToUpper(param.Unit().ID().GetText())
			if unit != "K" && unit != "M" {
				v.Errors += "MKDISK: Tipo de unidad inválido. Use K o M.\n"
				return nil
			}
		default:
			v.Errors += "MKDISK: Parámetro desconocido.\n"
			return nil
		}
	}

	// Validaciones obligatorias
	if size == 0 {
		v.Errors += "MKDISK: Falta el parámetro de 'size'."
		return nil
	}
	if path == "" {
		v.Errors += "MKDISK: Falta el parámetro de 'path'."
		return nil
	}

	// Crear el disco físico
	err := v.DiskManager.Mkdisk(size, fit, unit, path)
	if err != nil {
		v.Errors += fmt.Sprintf("MKDISK: %s", err.Error())
		return nil
	}

	fileName := filepath.Base(path) // Disco1.mia
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	v.Console += fmt.Sprintf("MKDISK: El disco %s se creó en %s correctamente.\n", name, path)
	return nil
}

// Eliminar un disco
func (v *Visitor) VisitRMDISK(ctx *parser.RMDISKContext) interface{} {
	path := ""

	// Solo hay un parámetro (path) en la regla rmdisk_param
	if ctx.Rmdisk_param() != nil && ctx.Rmdisk_param().Path() != nil {
		param := ctx.Rmdisk_param().Path()
		path = extractValue(param.STRING_LIT(), param.UNQUOTED_TEXT())
		if !strings.HasSuffix(strings.ToLower(path), ".mia") {
			path += ".mia"
		}
	}

	if path == "" {
		v.Errors += "RMDISK: Falta el parámetro de ruta.\n"
		return nil
	}

	err := v.DiskManager.Rmdisk(path)
	if err != nil {
		v.Errors += fmt.Sprintf("RMDISK: %s", err.Error())
		return nil
	}

	fileName := filepath.Base(path) // Disco1.mia
	name := cleanQuotes(strings.TrimSuffix(fileName, filepath.Ext(fileName)))

	v.Console += fmt.Sprintf("RMDISK: El disco %s con el path %s se eliminó correctamente.\n", name, path)
	return nil
}

// Crear partición en un disco
func (v *Visitor) VisitFDISK(ctx *parser.FDISKContext) interface{} {
	var size, add int64
	var path, name, delete string
	unit := "K"     // por defecto
	partType := "P" // primaria por defecto
	fit := "WF"     // worst fit por defecto

	// Para controlar duplicados
	used := make(map[string]bool)

	// Recorremos todos los parámetros
	for _, param := range ctx.Fdisk_params().AllFdisk_param() {
		switch {
		case param.Size() != nil:
			if used["size"] {
				v.Errors += "FDISK: El parámetro 'size' está duplicado.\n"
				return nil
			}
			used["size"] = true

			textSize := param.Size().INT_LIT().GetText()
			parsedSize, err := strconv.ParseInt(textSize, 10, 64)
			if err != nil || parsedSize <= 0 {
				v.Errors += "FDISK: El parámetro 'size' debe ser un número positivo.\n"
				return nil
			}
			size = parsedSize

		case param.Path() != nil:
			if used["path"] {
				v.Errors += "FDISK: El parámetro 'path' está duplicado.\n"
				return nil
			}
			used["path"] = true

			p := param.Path()
			path = extractValue(p.STRING_LIT(), p.UNQUOTED_TEXT())
			if !strings.HasSuffix(strings.ToLower(path), ".mia") {
				path += ".mia"
			}

		case param.Name() != nil:
			if used["name"] {
				v.Errors += "FDISK: El parámetro 'name' está duplicado.\n"
				return nil
			}
			used["name"] = true

			children := param.Name().GetChildren()
			if len(children) == 0 {
				v.Errors += "FDISK: El parámetro 'name' es inválido.\n"
				return nil
			}

			// Último token normalmente es el valor del nombre
			name = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

			if len(name) == 0 || len(name) > 16 {
				v.Errors += "FDISK: El parámetro 'name' debe tener entre 1 y 16 caracteres.\n"
				return nil
			}

		case param.Unit() != nil:
			if used["unit"] {
				v.Errors += "FDISK: El parámetro 'unit' está duplicado.\n"
				return nil
			}
			used["unit"] = true

			children := param.Unit().GetChildren()
			if len(children) == 0 {
				v.Errors += "FDISK: El parámetro 'unit' es inválido.\n"
				return nil
			}

			unit = strings.ToUpper(children[len(children)-1].(antlr.TerminalNode).GetText())

			if unit != "B" && unit != "K" && unit != "M" {
				v.Errors += "FDISK: El tipo de unidad es inválido. Use B, K o M.\n"
				return nil
			}

		case param.Type_() != nil:
			if used["type"] {
				v.Errors += "FDISK: El parámetro 'type' está duplicado.\n"
				return nil
			}
			used["type"] = true

			children := param.Type_().GetChildren()
			if len(children) == 0 {
				v.Errors += "FDISK: El parámetro 'type' es inválido.\n"
				return nil
			}

			// El último hijo normalmente es el valor: P, E o L
			partType = cleanQuotes(strings.ToUpper(children[len(children)-1].(antlr.TerminalNode).GetText()))

			if partType != "P" && partType != "E" && partType != "L" {
				v.Errors += "FDISK: El tipo de partición es inválido. Use P, E o L.\n"
				return nil
			}

		case param.Fit() != nil:
			if used["fit"] {
				v.Errors += "FDISK: El parámetro 'fit' está duplicado.\n"
				return nil
			}
			used["fit"] = true

			fit = cleanQuotes(strings.ToUpper(param.Fit().ID().GetText()))
			if fit != "BF" && fit != "FF" && fit != "WF" {
				v.Errors += "FDISK: El tipo de ajuste es inválido. Use BF, FF o WF.\n"
				return nil
			}
		case param.Delete_() != nil:
			if used["delete"] {
				v.Errors += "FDISK: El parámetro 'delete' está duplicado.\n"
				return nil
			}
			used["delete"] = true

			delete = cleanQuotes(strings.ToUpper(param.Delete_().ID().GetText()))
			if delete != "FAST" && delete != "FULL" {
				v.Errors += "FDISK: El tipo de eliminación es inválido. Use FAST o FULL.\n"
				return nil
			}
		case param.Add() != nil:
			if used["add"] {
				v.Errors += "FDISK: El parámetro 'add' está duplicado."
				return nil
			}
			used["add"] = true

			children := param.Add().GetChildren()
			if len(children) == 0 {
				v.Errors += "FDISK: El parámetro 'add' es inválido."
				return nil
			}

			add = int64(extractInt(children[len(children)-1]))
			if add <= 0 {
				v.Errors += "FDISK: El parámetro 'add' debe ser un número entero positivo."
			}

			txt := param.Add().GetText()

			//tomar el valor negativo. valor válido
			if strings.Contains(txt, "=-") {
				add = -add
			}

		}
	}

	// Validaciones obligatorias
	if size == 0 {
		v.Errors += "FDISK: El parámetro 'size' es obligatorio.\n"
		return nil
	}
	if path == "" {
		v.Errors += "FDISK: El parámetro 'path' es obligatorio.\n"
		return nil
	}
	if name == "" {
		v.Errors += "FDISK: El parámetro 'name' es obligatorio.\n"
		return nil
	}

	// Crear partición
	err := v.DiskManager.Fdisk(path, size, unit[0], partType[0], fit, delete, name, add)
	if err != nil {
		v.Errors += fmt.Sprintf("FDISK: Error al crear la partición %s: %v", name, err)
		return nil
	}

	var message_type string
	switch partType[0] {
	case 'P':
		message_type = "primaria"
	case 'E':
		message_type = "extendida"
	case 'L':
		message_type = "lógica"
	}

	v.Console += fmt.Sprintf("FDISK: Partición '%s' del tipo %s creada exitosamente en el disco %s\n", name, message_type, path)
	return nil
}

// Montar una partición
func (v *Visitor) VisitMOUNT(ctx *parser.MOUNTContext) interface{} {
	var path, name string
	used := make(map[string]bool)

	// Recorremos los parámetros
	for _, param := range ctx.Mount_params().AllMount_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "MOUNT: Parámetro 'path' duplicado.\n"
				return nil
			}
			used["path"] = true

			p := param.Path()
			path = extractValue(p.STRING_LIT(), p.UNQUOTED_TEXT())
			if !strings.HasSuffix(strings.ToLower(path), ".mia") {
				path += ".mia"
			}

		case param.Name() != nil:
			if used["name"] {
				v.Errors += "MOUNT: Parámetro 'name' duplicado.\n"
				return nil
			}
			used["name"] = true

			children := param.Name().GetChildren()
			if len(children) == 0 {
				v.Errors += "MOUNT: Parámetro 'name' inválido.\n"
				return nil
			}

			// Último token normalmente es el valor del nombre
			name = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

			if len(name) == 0 || len(name) > 16 {
				v.Errors += "MOUNT: El nombre debe tener entre 1 y 16 caracteres.\n"
				return nil
			}
		}
	}

	// Validaciones obligatorias
	if path == "" {
		v.Errors += "MOUNT: Parámetro 'path' obligatorio.\n"
		return nil
	}
	if name == "" {
		v.Errors += "MOUNT: Parámetro 'name' obligatorio.\n"
		return nil
	}

	// Verificar que exista el archivo
	if _, err := os.Stat(path); os.IsNotExist(err) {
		v.Errors += fmt.Sprintf("MOUNT: El archivo de disco %s no existe.", path)
		return nil
	}

	// Montar la partición
	resultList, err := v.DiskManager.Mount(*v.ListMountedPartitions, path, name)
	if err != nil {
		v.Errors += fmt.Sprintf("MOUNT: Error al montar la partición: %v", err)
		return nil
	}

	// Actualizar lista
	v.ListMountedPartitions = &resultList

	// Buscar partición recién montada
	mounted := v.ListMountedPartitions.GetPartitionByPathAndName(path, name)
	if mounted == nil {
		v.Errors += "MOUNT: La partición montada no fue encontrada después de montar.\n"
		return nil
	}

	// Revisar si ya tiene FS
	hasFS, err := mounted.HasFileSystem()
	if err != nil {
		v.Errors += fmt.Sprintf("MOUNT: Error al verificar el sistema de archivos: %v", err)
		return nil
	}

	if hasFS {
		fs := ext2.NewEXT2(*mounted)
		mounted.FileSystem = fs
		v.Console += fmt.Sprintf("MOUNT: La partición '%s' ya tiene un sistema de archivos EXT2 y ha sido inicializada.\n", name)
	} else {
		mounted.FileSystem = nil
		v.Console += fmt.Sprintf("MOUNT: La partición '%s' ha sido montada.\n", name)
	}
	return nil
}

// Desmontar una partición
func (v *Visitor) VisitUNMOUNT(ctx *parser.UNMOUNTContext) interface{} {
	var id string
	used := make(map[string]bool)
	// Recorremos los parámetros

	if used["id"] {
		v.Errors += "UNMOUNT: Parámetro 'id' duplicado.\n"
		return nil
	}
	used["id"] = true
	children := ctx.Unmount_param().Id_text().GetChildren()
	if len(children) == 0 {
		v.Errors += "UNMOUNT: Parámetro 'id' inválido.\n"
		return nil
	}
	// Último token normalmente es el valor del ID
	id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
	if len(id) != 4 {
		v.Errors += "UNMOUNT: El ID debe tener exactamente 4 caracteres.\n"
		return nil
	}

	// Validaciones obligatorias
	if id == "" {
		v.Errors += "UNMOUNT: Parámetro 'id' obligatorio.\n"
		return nil
	}
	// Verificar que la partición esté montada
	if v.ListMountedPartitions.GetPathById(id) == "" {
		v.Errors += fmt.Sprintf("UNMOUNT: No se encontró ninguna partición montada para el ID %s.", id)
		return nil
	}
	// Desmontar la partición
	if err := v.DiskManager.Unmount(*v.ListMountedPartitions, id); err != nil {
		v.Errors += fmt.Sprintf("UNMOUNT: Error al desmontar la partición con ID %s: %v", id, err)
		return nil
	}
	v.Console += fmt.Sprintf("UNMOUNT: La partición con ID %s ha sido desmontada correctamente.\n", id)
	return nil

}

// Listar particiones montadas
func (v *Visitor) VisitMOUNTED(ctx *parser.MOUNTEDContext) interface{} {
	if mountedInfo, err := v.DiskManager.Mounted(*v.ListMountedPartitions); err != nil {
		v.Errors += fmt.Sprintf("MOUNTED: Error listando particiones montadas: %v", err)
	} else {
		v.Console += mountedInfo
	}
	return nil
}

// Crear un sistema de archivos en una partición montada
func (v *Visitor) VisitMKFS(ctx *parser.MKFSContext) interface{} {
	var id, typeFs, fsName string
	used := make(map[string]bool)

	// Valores por defecto
	typeFs = "full"
	fsName = "2fs"

	// Recorremos parámetros
	for _, param := range ctx.Mkfs_params().AllMkfs_param() {
		switch {
		case param.Id_text() != nil:
			if used["id"] {
				v.Errors += "MKFS: Parámetro 'id' duplicado.\n"
				return nil
			}
			used["id"] = true

			children := param.Id_text().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKFS: Parámetro 'id' inválido.\n"
				return nil
			}

			// Último token normalmente es el valor del ID
			id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

			if len(id) != 4 {
				v.Errors += "MKFS: El ID debe tener exactamente 4 caracteres.\n"
				return nil
			}

		case param.Type_() != nil:
			if used["type"] {
				v.Errors += "MKFS: Parámetro 'type' duplicado.\n"
				return nil
			}
			used["type"] = true

			children := param.Type_().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKFS: Parámetro 'type' inválido.\n"
				return nil
			}

			// Último token normalmente es el valor: "full" o "fast"
			typeFs = cleanQuotes(strings.ToLower(children[len(children)-1].(antlr.TerminalNode).GetText()))

			if typeFs != "full" && typeFs != "fast" {
				v.Errors += "MKFS: Tipo inválido. Use 'full' o 'fast'.\n"
				return nil
			}
		case param.Fs() != nil:
			if used["fs"] {
				v.Errors += "MKFS: Parámetro 'fs' duplicado.\n"
				return nil
			}
			used["fs"] = true

			children := param.Fs().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKFS: Parámetro 'fs' inválido.\n"
				return nil
			}

			// Último token normalmente es el valor: "ext2" o "ext3"
			fsName = cleanQuotes(strings.ToLower(children[len(children)-1].(antlr.TerminalNode).GetText()))

			if fsName != "2fs" && fsName != "3fs" {
				v.Errors += "MKFS: Sistema de archivos inválido. Use '2fs'(EXT2) o '3fs'(EXT3).\n"
				return nil
			}
		default:
			v.Errors += "MKFS: Parámetro desconocido.\n"
		}
	}

	// Verificar que la partición esté montada
	if v.ListMountedPartitions.GetPathById(id) == "" {
		v.Errors += fmt.Sprintf("MKFS: No se encontró ninguna partición montada para el ID %s.", id)
		return nil
	}

	// Crear el sistema de archivos
	fs, err := v.NewFileSystem(fsName, id)
	if err != nil {
		v.Errors += fmt.Sprintf("MKFS: Se ha producido un error al crear el sistema de archivos: %v", err)
		return nil
	}

	if err := fs.Mkfs(id, typeFs); err != nil {
		v.Errors += fmt.Sprintf("MKFS: Se ha producido un error al formatear el sistema de archivos: %v", err)
		return nil
	}

	// Asociar el FS a la partición montada
	p := v.ListMountedPartitions.GetPartitionById(id)
	if p == nil {
		v.Errors += fmt.Sprintf("MKFS: No se encontró partición montada con id %s.\n", id)
		return nil
	}
	p.FileSystem = fs

	v.Console += fmt.Sprintf("MKFS: La partición con ID %s ha sido formateada correctamente con %s (%s).\n", id, fsName, typeFs)
	return nil
}

// Mostrar contenido de archivos
func (v *Visitor) VisitCAT(ctx *parser.CATContext) interface{} {
	type FileArg struct {
		Num  int
		Path string
	}

	var files []FileArg
	used := make(map[int]bool)

	for _, filen := range ctx.Cat_params().AllCat_param() {
		// Extraer texto del nombre (ej: "file1", "file2", etc.)
		paramName := strings.ToLower(filen.Filen().GetText())
		matches := fileRegex.FindStringSubmatch(paramName)
		if len(matches) < 2 {
			v.Errors += fmt.Sprintf("CAT: Parámetro inválido %q, se esperaba -fileN.\n", paramName)
			continue
		}

		// Convertir número
		n, err := strconv.Atoi(matches[1])
		if err != nil || n <= 0 {
			v.Errors += fmt.Sprintf("CAT: Número inválido en parámetro %q.\n", paramName)
			continue
		}
		if used[n] {
			v.Errors += fmt.Sprintf("CAT: Parámetro duplicado -file%d.\n", n)
			continue
		}
		used[n] = true

		// Obtener la ruta (STRING o TEXTO sin comillas)
		filePath := extractValue(filen.Filen().STRING_LIT(), filen.Filen().UNQUOTED_TEXT())
		if filePath == "" {
			v.Errors += fmt.Sprintf("CAT: Falta el 'path' para -file%d.\n", n)
			continue
		}

		files = append(files, FileArg{Num: n, Path: filePath})
	}

	if len(files) == 0 {
		return nil
	}

	// Ordenar por número de archivo
	sort.Slice(files, func(i, j int) bool {
		return files[i].Num < files[j].Num
	})

	// Validar sesión activa
	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "CAT: La partición con sesión activa no tiene un sistema de archivos.\n"
		return nil
	}

	// Recolectar paths
	var paths []string
	for _, f := range files {
		paths = append(paths, f.Path)
	}

	// Ejecutar CAT
	result, err := part.FileSystem.Cat(paths...)
	if err != nil {
		v.Errors += fmt.Sprintf("CAT: Error al ejecutar la acción: %v", err)
		return nil
	}

	// Imprimir y guardar en consola
	v.Console += result + "\n"
	return nil
}

// Iniciar sesión
func (v *Visitor) VisitLOGIN(ctx *parser.LOGINContext) interface{} {
	var user, pass, id string
	used := make(map[string]bool)

	for _, param := range ctx.Login_params().AllLogin_param() {
		switch {
		case param.User() != nil:
			if used["user"] {
				v.Errors += "LOGIN: El parámetro 'user' está duplicado."
				return nil
			}
			used["user"] = true

			children := param.User().GetChildren()
			if len(children) == 0 {
				v.Errors += "LOGIN: El parámetro 'user' es inválido."
				return nil
			}
			user = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

		case param.Pass() != nil:
			if used["pass"] {
				v.Errors += "LOGIN: El parámetro 'pass' está duplicado."
				return nil
			}
			used["pass"] = true

			children := param.Pass().GetChildren()
			if len(children) == 0 {
				v.Errors += "LOGIN: El parámetro 'pass' es inválido."
				return nil
			}
			pass = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

		case param.Id_text() != nil:
			if used["id"] {
				v.Errors += "LOGIN: El parámetro 'id' está duplicado."
				return nil
			}
			used["id"] = true

			children := param.Id_text().GetChildren()
			if len(children) == 0 {
				v.Errors += "LOGIN: El parámetro 'id' es inválido."
				return nil
			}
			id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
		}
	}

	// Validaciones obligatorias
	if user == "" {
		v.Errors += "LOGIN: Falta el parámetro 'user'."
	}
	if pass == "" {
		v.Errors += "LOGIN: Falta el parámetro 'pass'."
	}
	if id == "" {
		v.Errors += "LOGIN: Falta el parámetro 'id'."
	}
	if len(v.Errors) > 0 {
		return nil
	}

	// Validar si ya hay sesión activa
	if v.IdMountedAndLogued != "" {
		v.Errors += fmt.Sprintf("LOGIN: Ya hay una sesión activa en la partición %s", v.IdMountedAndLogued)
		return nil
	}

	// Validar partición montada
	part := v.ListMountedPartitions.GetPartitionById(strings.Trim(id, " "))
	if part == nil || part.FileSystem == nil {
		v.Errors += fmt.Sprintf("LOGIN: No se encontró una partición montada con ID %s o no tiene sistema de archivos.", id)
		return nil
	}

	// Intentar login
	if err := part.FileSystem.Login(user, pass); err != nil {
		v.Errors += fmt.Sprintf("LOGIN Al iniciar sesión: %v", err)
		return nil
	}

	v.IdMountedAndLogued = id
	v.Console += fmt.Sprintf("LOGIN: Logueado como %s en la partición %s\n", user, v.IdMountedAndLogued)
	return nil
}

// Cerrar sesión
func (v *Visitor) VisitLOGOUT(ctx *parser.LOGOUTContext) interface{} {
	if v.IdMountedAndLogued == "" {
		v.Errors += "LOGOUT: No hay una sesión activa para cerrar sesión."
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "LOGOUT: No se encontró un sistema de archivos para la sesión actual."
		v.IdMountedAndLogued = ""
		return nil
	}

	if err := part.FileSystem.Logout(); err != nil {
		v.Errors += fmt.Sprintf("LOGOUT: Error al cerrar sesión: %v", err)
		return nil
	}

	v.Console += "LOGOUT: Cierre de sesión exitoso!\n"
	v.IdMountedAndLogued = ""
	return nil
}

// Crear un grupo
func (v *Visitor) VisitMKGRP(ctx *parser.MKGRPContext) interface{} {
	var name string

	children := ctx.Mkgrp_param().Name().GetChildren()
	if len(children) == 0 {
		v.Errors += "MKGRP: El parámetro 'name' es inválido."
		return nil
	}

	// Último token normalmente es el valor del nombre
	name = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

	if len(name) == 0 || len(name) > 10 {
		v.Errors += "MKGRP: El nombre debe tener entre 1 y 10 caracteres."
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "MKGRP: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Mkgrp(name); err != nil {
		v.Errors += fmt.Sprintf("MKGRP Al crear el grupo '%s': %v", name, err)
	}

	v.Console += fmt.Sprintf("MKGRP: Grupo '%s' creado exitosamente.\n", name)

	return nil
}

// Eliminar un grupo
func (v *Visitor) VisitRMGRP(ctx *parser.RMGRPContext) interface{} {
	var name string

	children := ctx.Rmgrp_param().Name().GetChildren()
	if len(children) == 0 {
		v.Errors += "RMGRP: El parámetro 'name' es inválido."
		return nil
	}

	// Último token normalmente es el valor del nombre
	name = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
	if len(name) == 0 || len(name) > 10 {
		v.Errors += "RMGRP: El nombre debe tener entre 1 y 10 caracteres."
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "RMGRP: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Rmgrp(name); err != nil {
		v.Errors += fmt.Sprintf("RMGRP: Error al eliminar el grupo '%s': %v", name, err)
	}

	v.Console += fmt.Sprintf("RMGRP: Grupo '%s' eliminado exitosamente.\n", name)

	return nil
}

// Crear un usuario
func (v *Visitor) VisitMKUSR(ctx *parser.MKUSRContext) interface{} {
	var user, pass, grp string
	used := make(map[string]bool)

	for _, param := range ctx.Mkusr_params().AllMkusr_param() {
		switch {
		case param.User() != nil:
			if used["user"] {
				v.Errors += "MKUSR: Parámetro de 'user' duplicado."
				return nil
			}
			used["user"] = true

			children := param.User().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKUSR: El parámetro 'name' es inválido."
				return nil
			}

			user = cleanQuotes(strings.ToLower(children[len(children)-1].(antlr.TerminalNode).GetText()))

			if len(user) == 0 || len(user) > 10 {
				v.Errors += "MKUSR: El parámetro 'user' debe tener entre 1 y 10 caracteres."
			}

		case param.Pass() != nil:
			if used["pass"] {
				v.Errors += "MKUSR: Parámetro de 'pass' duplicado."
				return nil
			}
			used["pass"] = true

			children := param.Pass().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKUSR: El parámetro 'pass' es inválido."
				return nil
			}

			pass = cleanQuotes(strings.ToLower(children[len(children)-1].(antlr.TerminalNode).GetText()))

			if len(pass) == 0 || len(pass) > 10 {
				v.Errors += "MKUSR: El parámetro 'pass' debe tener entre 1 y 10 caracteres."
			}

		case param.Grp() != nil:
			if used["grp"] {
				v.Errors += "MKUSR: Parámetro de 'grp' duplicado."
				return nil
			}
			used["grp"] = true

			children := param.Grp().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKUSR: El parámetro 'grp' es inválido."
				return nil
			}

			grp = cleanQuotes(strings.ToLower((children[len(children)-1].(antlr.TerminalNode).GetText())))

			if len(grp) == 0 || len(grp) > 10 {
				v.Errors += "MKUSR: El parámetro 'grp' debe tener entre 1 y 10 caracteres."
			}
		}
	}

	if user == "" {
		v.Errors += "MKUSR: Falta el parámetro 'user'."
	}
	if pass == "" {
		v.Errors += "MKUSR: Falta el parámetro 'pass'."
	}
	if grp == "" {
		v.Errors += "MKUSR: Falta el parámetro 'grp'."
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "MKUSR: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Mkusr(user, pass, grp); err != nil {
		v.Errors += fmt.Sprintf("MKUSR: Error al crear el usuario '%s': %v", user, err)
	}

	v.Console += fmt.Sprintf("MKUSR: Usuario '%s' creado exitosamente.\n", user)

	return nil
}

// Eliminar un usuario
func (v *Visitor) VisitRMUSR(ctx *parser.RMUSRContext) interface{} {
	children := ctx.Rmusr_param().User().GetChildren()

	if len(children) == 0 {
		v.Errors += "RMUSR: El parámetro 'user' es inválido."
		return nil
	}
	user := cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())

	if len(user) == 0 || len(user) > 10 {
		v.Errors += "RMUSR: El parámetro 'user' debe tener entre 1 y 10 caracteres."
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "RMUSR: No hay un sistema de archivos encontrado."
		return nil
	}
	fs := part.FileSystem

	if err := fs.Rmusr(user); err != nil {
		v.Errors += fmt.Sprintf("RMUSR: Error al eliminar el usuario '%s': %v", user, err)
	}

	v.Console += fmt.Sprintf("RMUSR: Usuario '%s' eliminado exitosamente.\n", user)

	return nil
}

// Cambiar grupo de un usuario
func (v *Visitor) VisitCHGRP(ctx *parser.CHGRPContext) interface{} {
	var user, grp string
	used := make(map[string]bool)

	for _, param := range ctx.Chgrp_params().AllChgrp_param() {
		switch {
		case param.User() != nil:
			if used["user"] {
				v.Errors += "CHGRP: Parámetro de 'user' duplicado."
				return nil
			}
			used["user"] = true

			children := param.User().GetChildren()

			if len(children) == 0 {
				v.Errors += "CHGRP: El parámetro 'user' es inválido."
				return nil
			}
			user = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
			if len(user) == 0 || len(user) > 10 {
				v.Errors += "CHGRP: El parámetro 'user' debe tener entre 1 y 10 caracteres."
			}

		case param.Grp() != nil:
			if used["grp"] {
				v.Errors += "CHGRP: Parámetro de 'grp' duplicado."
				return nil
			}
			used["grp"] = true

			children := param.Grp().GetChildren()
			if len(children) == 0 {
				v.Errors += "CHGRP: El parámetro 'grp' es inválido."
				return nil
			}
			grp = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
			if len(grp) == 0 || len(grp) > 10 {
				v.Errors += "CHGRP: El parámetro 'grp' debe tener entre 1 y 10 caracteres."
			}
		}
	}

	if user == "" {
		v.Errors += "CHGRP: Falta el parámetro 'user'."
	}
	if grp == "" {
		v.Errors += "CHGRP: Falta el parámetro 'grp'."
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "CHGRP: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Chgrp(user, grp); err != nil {
		v.Errors += fmt.Sprintf("CHGRP: Error al cambiar el grupo del usuario '%s': %v", user, err)
	}

	v.Console += fmt.Sprintf("CHGRP: Grupo del usuario '%s' cambiado a '%s' exitosamente.\n", user, grp)

	return nil
}

// Crear un archivo
func (v *Visitor) VisitMKFILE(ctx *parser.MKFILEContext) interface{} {
	var path, cont string
	createDir := false
	var sizeInt int
	used := make(map[string]bool)

	for _, param := range ctx.Mkfile_params().AllMkfile_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "MKFILE: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.R() != nil:
			if used["r"] {
				v.Errors += "MKFILE: El parámetro 'r' está duplicado."
				return nil
			}
			used["r"] = true
			createDir = true

		case param.Size() != nil:
			if used["size"] {
				v.Errors += "MKFILE: El parámetro 'size' está duplicado."
				return nil
			}
			used["size"] = true

			txt := param.Size().GetText()

			if strings.Contains(txt, "=-") {
				v.Errors += "MKFILE: El parámetro 'size' no puede ser negativo."
				return nil
			}

			children := param.Size().GetChildren()
			if len(children) == 0 {
				v.Errors += "MKFILE: El parámetro 'size' es inválido."
				return nil
			}

			sizeInt = extractInt(children[len(children)-1])
			if sizeInt <= 0 {
				v.Errors += "MKFILE: El parámetro 'size' debe ser un número entero positivo."
			}

		case param.Cont() != nil:
			if used["cont"] {
				v.Errors += "MKFILE: El parámetro 'cont' está duplicado."
				return nil
			}
			used["cont"] = true
			cont = extractValue(param.Cont().STRING_LIT(), param.Cont().UNQUOTED_TEXT())
		}
	}

	if path == "" {
		v.Errors += "MKFILE: Falta el parámetro 'path'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "MKFILE: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Mkfile(path, createDir, int64(sizeInt), cont); err != nil {
		v.Errors += fmt.Sprintf("MKFILE: Error creando el archivo '%s': %v", path, err)
	}

	v.Console += fmt.Sprintf("MKFILE: Archivo '%s' creado exitosamente.\n", path)

	return nil
}

// Crear un directorio
func (v *Visitor) VisitMKDIR(ctx *parser.MKDIRContext) interface{} {
	var path string
	pFlag := false
	used := make(map[string]bool)

	for _, param := range ctx.Mkdir_params().AllMkdir_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "MKDIR: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.P() != nil:
			if used["p"] {
				v.Errors += "MKDIR: El parámetro 'p' está duplicado."
				return nil
			}
			used["p"] = true
			pFlag = true
		}
	}

	if path == "" {
		v.Errors += "MKDIR: Falta el parámetro 'path'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "MKDIR: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Mkdir(path, pFlag); err != nil {
		v.Errors += fmt.Sprintf("MKDIR: Error creando el directorio '%s': %v", path, err)
	}

	v.Console += fmt.Sprintf("MKDIR: Directorio '%s' creado exitosamente.\n", path)

	return nil
}

// Eliminar un archivo o directorio
func (v *Visitor) VisitREMOVE(ctx *parser.REMOVEContext) interface{} {
	var path string

	used := make(map[string]bool)

	if used["path"] {
		v.Errors += "REMOVE: El parámetro 'path' está duplicado."
		return nil
	}
	used["path"] = true
	path = strings.Trim(extractValue(ctx.Remove_param().Path().STRING_LIT(), ctx.Remove_param().Path().UNQUOTED_TEXT()), "\"")

	if path == "" {
		v.Errors += "REMOVE: Falta el parámetro 'path'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "REMOVE: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Remove(path); err != nil {
		v.Errors += fmt.Sprintf("REMOVE: Error eliminando el archivo o directorio '%s': %v", path, err)
		return nil
	}

	v.Console += fmt.Sprintf("REMOVE: Archivo o directorio '%s' eliminado exitosamente.\n", path)

	return nil
}

// Editar un archivo
func (v *Visitor) VisitEDIT(ctx *parser.EDITContext) interface{} {
	var path, contenido string
	used := make(map[string]bool)

	for _, param := range ctx.Edit_params().AllEdit_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "EDIT: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Contenido() != nil:
			if used["contenido"] {
				v.Errors += "EDIT: El parámetro 'contenido' está duplicado."
				return nil
			}
			used["contenido"] = true
			contenido = extractValue(param.Contenido().STRING_LIT(), param.Contenido().UNQUOTED_TEXT())
		}
	}
	if path == "" {
		v.Errors += "EDIT: Falta el parámetro 'path'."
		return nil
	}
	if contenido == "" {
		v.Errors += "EDIT: Falta el parámetro 'contenido'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "EDIT: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Edit(path, contenido); err != nil {
		v.Errors += fmt.Sprintf("EDIT: Error editando el archivo '%s': %v", path, err)
		return nil
	}

	v.Console += fmt.Sprintf("EDIT: Archivo '%s' editado exitosamente.\n", path)

	return nil
}

// Renombrar un archivo o directorio
func (v *Visitor) VisitRENAME(ctx *parser.RENAMEContext) interface{} {
	var path, name string
	used := make(map[string]bool)

	for _, param := range ctx.Rename_params().AllRename_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "RENAME: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Name() != nil:
			if used["name"] {
				v.Errors += "RENAME: El parámetro 'name' está duplicado."
				return nil
			}
			used["name"] = true
			name = cleanQuotes(param.Name().GetText())
		}
	}
	if path == "" {
		v.Errors += "RENAME: Falta el parámetro 'path'."
		return nil
	}
	if name == "" {
		v.Errors += "RENAME: Falta el parámetro 'name'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)

	if part == nil || part.FileSystem == nil {
		v.Errors += "RENAME: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Rename(path, name); err != nil {
		v.Errors += fmt.Sprintf("RENAME: Error renombrando el archivo o directorio '%s': %v", path, err)
		return nil
	}

	v.Console += fmt.Sprintf("RENAME: Archivo o directorio '%s' renombrado a '%s' exitosamente.\n", path, name)

	return nil
}

// Copiar un archivo o directorio
func (v *Visitor) VisitCOPY(ctx *parser.COPYContext) interface{} {
	var path, destino string
	used := make(map[string]bool)

	for _, param := range ctx.Copy_params().AllCopy_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "COPY: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Destino() != nil:
			if used["destino"] {
				v.Errors += "COPY: El parámetro 'destino' está duplicado."
				return nil
			}
			used["destino"] = true
			destino = strings.Trim(extractValue(param.Destino().STRING_LIT(), param.Destino().UNQUOTED_TEXT()), "\"")
		}
	}
	if path == "" {
		v.Errors += "COPY: Falta el parámetro 'path'."
		return nil
	}
	if destino == "" {
		v.Errors += "COPY: Falta el parámetro 'destino'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "COPY: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Copy(path, destino); err != nil {
		v.Errors += fmt.Sprintf("COPY: Error copiando de '%s' a '%s': %v", path, destino, err)
		return nil
	}

	v.Console += fmt.Sprintf("COPY: Copiado de '%s' a '%s' exitosamente.\n", path, destino)

	return nil
}

// Mover un archivo o directorio
func (v *Visitor) VisitMOVE(ctx *parser.MOVEContext) interface{} {
	var path, destino string
	used := make(map[string]bool)

	for _, param := range ctx.Move_params().AllMove_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "MOVE: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Destino() != nil:
			if used["destino"] {
				v.Errors += "MOVE: El parámetro 'destino' está duplicado."
				return nil
			}
			used["destino"] = true
			destino = strings.Trim(extractValue(param.Destino().STRING_LIT(), param.Destino().UNQUOTED_TEXT()), "\"")
		}
	}
	if path == "" {
		v.Errors += "MOVE: Falta el parámetro 'path'."
		return nil
	}
	if destino == "" {
		v.Errors += "MOVE: Falta el parámetro 'destino'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "MOVE: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Move(path, destino); err != nil {
		v.Errors += fmt.Sprintf("MOVE: Error moviendo de '%s' a '%s': %v", path, destino, err)
		return nil
	}

	v.Console += fmt.Sprintf("MOVE: Movido de '%s' a '%s' exitosamente.\n", path, destino)

	return nil
}

// Encontrar un archivo o directorio
func (v *Visitor) VisitFIND(ctx *parser.FINDContext) interface{} {
	var path, name string
	used := make(map[string]bool)

	for _, param := range ctx.Find_params().AllFind_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "FIND: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Name_find() != nil:
			if used["name"] {
				v.Errors += "FIND: El parámetro 'name' está duplicado."
				return nil
			}
			used["name"] = true

			txt := param.Name_find().GetText()

			if strings.Contains(txt, "*") {
				name = "*"
			} else if strings.Contains(txt, "?") {
				name = "?"
			} else {
				name = strings.Trim(extractValue(param.Name_find().STRING_LIT(), param.Name_find().UNQUOTED_TEXT()), "\"")
			}
		}
	}

	if path == "" {
		v.Errors += "FIND: Falta el parámetro 'path'."
		return nil
	}
	if name == "" {
		v.Errors += "FIND: Falta el parámetro 'name'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "FIND: No hay un sistema de archivos encontrado."
		return nil
	}

	if find, err := part.FileSystem.Find(path, name); err != nil {
		v.Errors += fmt.Sprintf("FIND: Error encontrando de '%s' a '%s': %v", path, name, err)
		return nil
	} else if find == "" {
		v.Errors += fmt.Sprintf("FIND: No se encontró '%s' en '%s'.", name, path)
		return nil
	} else {
		//fmt.Println("Encontrado:", find)
	}

	v.Console += fmt.Sprintf("FIND: Encontrado de '%s' a '%s' exitosamente.\n", path, name)

	return nil
}

// Cambiar propietario o permisos de un archivo o directorio
func (v *Visitor) VisitCHOWN(ctx *parser.CHOWNContext) interface{} {

	//obligatorios: path, usuario
	//opcionales: -r
	var path, usuario string
	recursive := false
	used := make(map[string]bool)

	for _, param := range ctx.Chown_params().AllChown_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "CHOWN: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Usuario() != nil:
			if used["usuario"] {
				v.Errors += "CHOWN: El parámetro 'usuario' está duplicado."
				return nil
			}
			used["usuario"] = true

			children := param.Usuario().GetChildren()
			if len(children) == 0 {
				v.Errors += "CHOWN: El parámetro 'usuario' es inválido."
				return nil
			}
			usuario = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
		case param.R() != nil:
			if used["recursive"] {
				v.Errors += "CHOWN: El parámetro 'r' está duplicado."
				return nil
			}
			used["recursive"] = true
			recursive = true
		}
	}
	if path == "" {
		v.Errors += "CHOWN: Falta el parámetro 'path'."
		return nil
	}
	if usuario == "" {
		v.Errors += "CHOWN: Falta el parámetro 'usuario'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "CHOWN: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Chown(path, usuario, recursive); err != nil {
		v.Errors += fmt.Sprintf("CHOWN: Error cambiando propietario de '%s' a '%s': %v", path, usuario, err)
		return nil
	}

	v.Console += fmt.Sprintf("CHOWN: Propietario de '%s' cambiado a '%s' exitosamente.\n", path, usuario)

	return nil
}

// Cambiar permisos de un archivo o directorio
func (v *Visitor) VisitCHMOD(ctx *parser.CHMODContext) interface{} {
	var path string
	var ugo int64
	recursive := false

	used := make(map[string]bool)

	for _, param := range ctx.Chmod_params().AllChmod_param() {
		switch {
		case param.Path() != nil:
			if used["path"] {
				v.Errors += "CHMOD: El parámetro 'path' está duplicado."
				return nil
			}
			used["path"] = true
			path = strings.Trim(extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT()), "\"")

		case param.Ugo() != nil:
			if used["ugo"] {
				v.Errors += "CHMOD: El parámetro 'ugo' está duplicado."
				return nil
			}
			used["ugo"] = true

			children := param.Ugo().GetChildren()
			if len(children) == 0 {
				v.Errors += "CHMOD: El parámetro 'ugo' es inválido."
				return nil
			}

			ugo = int64(extractInt(children[len(children)-1]))
			if ugo < 0 || ugo > 777 {
				v.Errors += "CHMOD: El parámetro 'ugo' debe estar entre 0 y 777."
			}

		case param.R() != nil:
			if used["recursive"] {
				v.Errors += "CHMOD: El parámetro 'r' está duplicado."
				return nil
			}
			used["recursive"] = true
			recursive = true
		}
	}
	if path == "" {
		v.Errors += "CHMOD: Falta el parámetro 'path'."
		return nil
	}
	if ugo == 0 {
		v.Errors += "CHMOD: Falta el parámetro 'ugo'."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(v.IdMountedAndLogued)
	if part == nil || part.FileSystem == nil {
		v.Errors += "CHMOD: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Chmod(path, ugo, recursive); err != nil {
		v.Errors += fmt.Sprintf("CHMOD: Error cambiando permisos de '%s' a '%d': %v", path, ugo, err)
		return nil
	}

	v.Console += fmt.Sprintf("CHMOD: Permisos de '%s' cambiados a '%d' exitosamente.\n", path, ugo)

	return nil
}

// Recuperar espacio de un archivo o directorio
func (v *Visitor) VisitRECOVERY(ctx *parser.RECOVERYContext) interface{} {
	var id string

	used := make(map[string]bool)
	if used["id"] {
		v.Errors += "RECOVERY: El parámetro 'id' está duplicado."
		return nil
	}
	used["id"] = true

	children := ctx.Recovery_param().Id_text().GetChildren()
	if len(children) == 0 {
		v.Errors += "RECOVERY: El parámetro 'id' es inválido."
		return nil
	}

	id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
	if len(id) != 4 {
		v.Errors += "RECOVERY: El parámetro 'id' debe tener exactamente 4 caracteres."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(id)
	if part == nil || part.FileSystem == nil {
		v.Errors += "RECOVERY: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Recovery(); err != nil {
		v.Errors += fmt.Sprintf("RECOVERY: Error recuperando espacio: %v", err)
		return nil
	}

	v.Console += "RECOVERY: Espacio recuperado exitosamente.\n"

	return nil
}

// Emular una perdida de datos
func (v *Visitor) VisitLOSS(ctx *parser.LOSSContext) interface{} {
	var id string

	used := make(map[string]bool)
	if used["id"] {
		v.Errors += "LOSS: El parámetro 'id' está duplicado."
		return nil
	}
	used["id"] = true

	children := ctx.Loss_param().Id_text().GetChildren()
	if len(children) == 0 {
		v.Errors += "LOSS: El parámetro 'id' es inválido."
		return nil
	}

	id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
	if len(id) != 4 {
		v.Errors += "LOSS: El parámetro 'id' debe tener exactamente 4 caracteres."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(id)
	if part == nil || part.FileSystem == nil {
		v.Errors += "LOSS: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Loss(); err != nil {
		v.Errors += fmt.Sprintf("LOSS: Error recuperando espacio: %v", err)
		return nil
	}

	v.Console += "LOSS: Espacio recuperado exitosamente.\n"

	return nil
}

// Emular una journaling
func (v *Visitor) VisitJOURNALING(ctx *parser.JOURNALINGContext) interface{} {
	var id string

	used := make(map[string]bool)
	if used["id"] {
		v.Errors += "JOURNALING: El parámetro 'id' está duplicado."
		return nil
	}
	used["id"] = true

	children := ctx.Journaling_param().Id_text().GetChildren()
	if len(children) == 0 {
		v.Errors += "JOURNALING: El parámetro 'id' es inválido."
		return nil
	}

	id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
	if len(id) != 4 {
		v.Errors += "JOURNALING: El parámetro 'id' debe tener exactamente 4 caracteres."
		return nil
	}
	if len(v.Errors) > 0 {
		return nil
	}

	part := v.ListMountedPartitions.GetPartitionById(id)
	if part == nil || part.FileSystem == nil {
		v.Errors += "JOURNALING: No hay un sistema de archivos encontrado."
		return nil
	}

	if err := part.FileSystem.Journaling(); err != nil {
		v.Errors += fmt.Sprintf("JOURNALING: Error recuperando espacio: %v", err)
		return nil
	}

	v.Console += "JOURNALING: Espacio recuperado exitosamente.\n"

	return nil
}

// Crear un reporte
func (v *Visitor) VisitREP(ctx *parser.REPContext) interface{} {
	var name, path, id, pathFileLS string
	used := make(map[string]bool)

	// --- Parsear parámetros ---
	for _, param := range ctx.Rep_params().AllRep_param() {
		switch {
		case param.Name() != nil:
			if used["name"] {
				v.Errors += "REP: Parámetro 'name' duplicado.\n"
				return nil
			}
			used["name"] = true

			children := param.Name().GetChildren()
			if len(children) == 0 {
				v.Errors += "REP: Parámetro 'name' inválido.\n"
				return nil
			}

			name = strings.ToLower(cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText()))

			switch name {
			case "mbr", "disk", "inode", "block", "bm_inode", "bm_block", "tree", "sb", "file", "ls":
			default:
				v.Errors += fmt.Sprintf("REP: Nombre de reporte inválido '%s'.\n", name)
				return nil
			}

		case param.Path() != nil:
			if used["path"] {
				v.Errors += "REP: Parámetro 'path' duplicado.\n"
				return nil
			}
			used["path"] = true
			path = extractValue(param.Path().STRING_LIT(), param.Path().UNQUOTED_TEXT())

		case param.Id_text() != nil:
			if used["id"] {
				v.Errors += "REP: Parámetro 'id' duplicado.\n"
				return nil
			}
			used["id"] = true

			children := param.Id_text().GetChildren()
			if len(children) == 0 {
				v.Errors += "REP: Parámetro 'id' inválido.\n"
				return nil
			}

			id = cleanQuotes(children[len(children)-1].(antlr.TerminalNode).GetText())
			if len(id) != 4 {
				v.Errors += "REP: El parámetro 'id' debe tener exactamente 4 caracteres.\n"
				return nil
			}

		case param.Path_file_ls() != nil:
			if used["path_file_ls"] {
				v.Errors += "REP: Parámetro 'path_file_ls' duplicado.\n"
				return nil
			}
			used["path_file_ls"] = true
			pathFileLS = extractValue(param.Path_file_ls().STRING_LIT(), param.Path_file_ls().UNQUOTED_TEXT())
		}
	}

	// --- Validaciones obligatorias ---
	if name == "" {
		v.Errors += "REP: Parámetro 'name' faltante.\n"
	}
	if path == "" {
		v.Errors += "REP: Parámetro 'path' faltante.\n"
	}
	if id == "" {
		v.Errors += "REP: Parámetro 'id' faltante.\n"
	}
	if (name == "file" || name == "ls") && pathFileLS == "" {
		v.Errors += "REP: Parámetro 'path_file_ls' faltante.\n"
	}
	if len(v.Errors) > 0 {
		return nil
	}

	// --- Extraer extensión ---
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
	if ext == "" {
		v.Errors += "REP: Extensión no definida.\n"
		return nil
	}

	// --- Validar partición montada ---
	partition := v.ListMountedPartitions.GetPartitionById(id)
	if partition == nil {
		v.Errors += fmt.Sprintf("REP: No se encontró ninguna partición montada para el ID '%s'.\n", id)
		return nil
	}

	// --- Reportes que no requieren FileSystem ---
	if name == "mbr" || name == "disk" {
		if err := v.DiskManager.Rep(*partition, name, path, id, pathFileLS, ext); err != nil {
			v.Errors += fmt.Sprintf("REP: Error generando reporte '%s': %v\n", name, err)
			return nil
		}
		v.Console += fmt.Sprintf("REP: Reporte '%s' generado exitosamente.\n", name)
		return nil
	}

	// --- Reportes que requieren FileSystem ---
	if partition.FileSystem == nil {
		v.Errors += fmt.Sprintf("REP: El reporte '%s' requiere un sistema de archivos montado en la partición '%s'.\n", name, id)
		return nil
	}

	if err := partition.FileSystem.Rep(name, path, pathFileLS, ext); err != nil {
		v.Errors += fmt.Sprintf("REP: Error generando reporte '%s': %v\n", name, err)
		return nil
	}

	// --- Éxito ---
	v.Console += fmt.Sprintf("REP: Reporte '%s' generado exitosamente.\n", name)
	return nil
}
