package dmanager

import (
	disk "backend/disk_structs"
	rep "backend/report_manager"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Constantes
const (
	STATUS_FREE = 0
	STATUS_USED = 1
)

// dependencia a partition

type DiskManager struct {
	FreeSpaceManager *FreeSpaceManager
	MbrManager       *MbrManager
	EbrManager       *EbrManager
}

func NewDiskManager() *DiskManager {
	return &DiskManager{
		FreeSpaceManager: &FreeSpaceManager{},
		MbrManager:       &MbrManager{},
		EbrManager:       &EbrManager{},
	}
}

func (d *DiskManager) CreatePrimary(name string, size int64, fit string, path string) error {

	// 1. Validaciones
	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("CreatePrimary: error al leer el MBR: %v", err)
	}

	partitionCount := 0
	for i := 0; i < 4; i++ {
		if mbr.Mbr_partition[i].Part_status == STATUS_USED {
			partitionCount++
			if strings.Trim(string(mbr.Mbr_partition[i].Part_name[:]), "\x00") == name {
				return fmt.Errorf("CreatePrimary: ya existe una partición con el nombre '%s'", name)
			}
		}
	}
	if partitionCount >= 4 {
		return fmt.Errorf("CreatePrimary: ya existen 4 particiones, no se pueden crear más")
	}

	// 2. Encontrar un hueco libre
	free := d.FreeSpaceManager.GetFreeSpaces(&mbr)

	var bestFitStart int64 = -1
	switch strings.ToUpper(fit) {
	case "FF":
		bestFitStart = d.FreeSpaceManager.FindFirstFit(size, free)
	case "BF":
		bestFitStart = d.FreeSpaceManager.FindBestFit(size, free)
	case "WF":
		bestFitStart = d.FreeSpaceManager.FindWorstFit(size, free)
	}
	if bestFitStart < 0 {
		return fmt.Errorf("CreatePrimary: no hay espacio suficiente para la partición primaria")
	}

	// 3. Crear la nueva estructura de Partición
	var newPartition disk.Partition
	newPartition.Part_status = '1' // ACTIVA
	newPartition.Part_type = 'P'
	newPartition.Part_fit = byte(strings.ToUpper(fit)[0])
	newPartition.Part_start = bestFitStart
	newPartition.Part_size = size
	copy(newPartition.Part_name[:], name)

	// 4. Añadirla a un slot vacío en el MBR
	added := false
	for i := 0; i < 4; i++ {
		if mbr.Mbr_partition[i].Part_status == STATUS_FREE { // libre
			mbr.Mbr_partition[i] = newPartition
			added = true
			break
		}
	}
	if !added {
		return fmt.Errorf("CreatePrimary: no se pudo añadir la partición al MBR")
	}

	// 5. Escribir el MBR actualizado de vuelta al disco
	err = d.MbrManager.WriteMBR(&mbr, path)
	if err != nil {
		return fmt.Errorf("CreatePrimary: error al escribir el MBR: %v", err)
	}

	return nil
}

func (d *DiskManager) CreateExtended(name string, size int64, fit string, path string) error {

	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("CreateExtended: error al leer el MBR: %v", err)
	}

	partitionCount := 0
	hasExtended := false
	for i := 0; i < 4; i++ {
		if mbr.Mbr_partition[i].Part_status == STATUS_USED {
			partitionCount++
			if mbr.Mbr_partition[i].Part_type == 'E' {
				hasExtended = true
			}
			if strings.Trim(string(mbr.Mbr_partition[i].Part_name[:]), "\x00") == name {
				return fmt.Errorf("CreateExtended: ya existe una partición con el nombre '%s'", name)
			}
		}
	}
	if partitionCount >= 4 {
		return fmt.Errorf("CreateExtended: ya existen 4 particiones, no se pueden crear más")
	}
	if hasExtended {
		return fmt.Errorf("CreateExtended: ya existe una partición extendida en este disco")
	}

	// 2. Buscar hueco disponible
	freeSpaces := d.FreeSpaceManager.GetFreeSpaces(&mbr)
	var bestFitStart int64 = -1

	switch strings.ToUpper(fit) {
	case "FF":
		bestFitStart = d.FreeSpaceManager.FindFirstFit(size, freeSpaces)
	case "BF":
		bestFitStart = d.FreeSpaceManager.FindBestFit(size, freeSpaces)
	case "WF":
		bestFitStart = d.FreeSpaceManager.FindWorstFit(size, freeSpaces)
	}

	if bestFitStart == -1 {
		return fmt.Errorf("CreateExtended: no hay suficiente espacio contiguo para la partición")
	}

	// 3. Verificar solapamiento
	for i := 0; i < 4; i++ {
		p := mbr.Mbr_partition[i]
		if p.Part_status == STATUS_USED {
			end := p.Part_start + p.Part_size
			if !(bestFitStart+size <= p.Part_start || bestFitStart >= end) {
				return fmt.Errorf("CreateExtended:la nueva partición se solapa con '%s'", string(p.Part_name[:]))
			}
		}
	}

	// 4. Crear la partición extendida
	var newPartition disk.Partition
	newPartition.Part_status = STATUS_USED
	newPartition.Part_type = 'E'
	newPartition.Part_fit = byte(strings.ToUpper(fit)[0])
	newPartition.Part_start = bestFitStart
	newPartition.Part_size = size
	copy(newPartition.Part_name[:], name)

	// 5. Guardarla en el MBR
	added := false
	for i := 0; i < 4; i++ {
		if mbr.Mbr_partition[i].Part_status == STATUS_FREE {
			mbr.Mbr_partition[i] = newPartition
			added = true
			break
		}
	}
	if !added {
		return fmt.Errorf("CreateExtended: no se encontró un slot libre en el MBR")
	}

	// 6. Escribir MBR en disco
	err = d.MbrManager.WriteMBR(&mbr, path)
	if err != nil {
		return fmt.Errorf("CreateExtended: error al escribir el MBR: %v", err)
	}

	// 7. Inicializar primer EBR
	var firstEBR disk.EBR
	firstEBR.Part_mount = STATUS_FREE
	firstEBR.Part_fit = newPartition.Part_fit
	firstEBR.Part_start = newPartition.Part_start
	firstEBR.Part_size = 0
	firstEBR.Part_next = -1

	err = d.EbrManager.WriteEBR(&firstEBR, path, newPartition.Part_start)
	if err != nil {
		return fmt.Errorf("CreateExtended: error al inicializar el primer EBR: %v", err)
	}

	return nil
}

func (d *DiskManager) CreateLogical(name string, size int64, fit string, path string) error {
	// 1. Buscar la partición extendida en el MBR
	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("CreateLogical: error al leer el MBR: %v", err)
	}

	var extendedPartition *disk.Partition
	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		if p.Part_status == STATUS_USED && p.Part_type == 'E' {
			extendedPartition = p
			break
		}
	}

	if extendedPartition == nil {
		return fmt.Errorf("CreateLogical: no existe partición extendida en el disco, no se puede crear lógica")
	}

	// 2. Cargar la lista de EBRs desde disco
	ebrs, err := d.EbrManager.GetEBRs(path, *extendedPartition)
	if err != nil {
		return fmt.Errorf("CreateLogical: error al cargar los EBRs: %v", err)
	}

	// 3. Validar que no exista una partición lógica con el mismo nombre
	for _, e := range ebrs {
		existingName := strings.Trim(string(e.Part_name[:]), "\x00")
		if e.Part_mount == STATUS_USED && existingName == name {
			return fmt.Errorf("CreateLogical: ya existe una partición lógica con el nombre '%s'", name)
		}
	}

	// 4. Buscar espacio libre dentro de la extendida
	free := d.FreeSpaceManager.GetFreeSpacesInExtended(*extendedPartition, ebrs)

	var start int64 = -1
	switch strings.ToUpper(fit) {
	case "FF":
		start = d.FreeSpaceManager.FindFirstFit(size, free)
	case "BF":
		start = d.FreeSpaceManager.FindBestFit(size, free)
	case "WF":
		start = d.FreeSpaceManager.FindWorstFit(size, free)
	default:
		return fmt.Errorf("CreateLogical: ajuste '%s' no soportado para lógica", fit)
	}

	if start < 0 {
		return fmt.Errorf("CreateLogical: no hay espacio suficiente en la extendida para la lógica")
	}

	// 5. Crear el nuevo EBR
	var newEbr disk.EBR
	newEbr.Part_mount = STATUS_USED
	newEbr.Part_fit = byte(strings.ToUpper(fit)[0])
	newEbr.Part_start = start
	newEbr.Part_size = size
	newEbr.Part_next = -1
	copy(newEbr.Part_name[:], name)

	// 6. Insertar en la lista de EBRs y enlazar
	err = d.EbrManager.AddEBR(path, extendedPartition, &newEbr, ebrs)
	if err != nil {
		return fmt.Errorf("CreateLogical: error al insertar el nuevo EBR: %v", err)
	}

	return nil
}

func (d *DiskManager) Mkdisk(size int64, fit string, unit string, path string) error {

	// Calcular tamaño en bytes
	if unit == "K" {
		size = size * 1024
	} else if unit == "M" {
		size = size * 1024 * 1024
	}

	fechaStr := time.Now().Format("2006-01-02 15:04:05")
	var fecha [20]byte
	copy(fecha[:], fechaStr)
	diskSignature := rand.Int63()

	// Crear el MBR
	mbr := disk.MBR{
		Mbr_tamano:         size,
		Mbr_fecha_creacion: fecha,
		Mbr_disk_signature: diskSignature,
		Dsk_fit:            fit[0],
		Mbr_partition:      [4]disk.Partition{},
	}

	if !strings.HasSuffix(strings.ToLower(path), ".mia") {
		path += ".mia"
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Mkdisk: error al crear directorios: %v", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Mkdisk: error al crear el archivo: %v", err)
	}
	defer file.Close()

	// Crear el archivo con el tamaño solicitado
	if err := file.Truncate(mbr.Mbr_tamano); err != nil {
		return fmt.Errorf("Mkdisk: error al truncar el archivo: %v", err)
	}

	// Escribir el MBR en los primeros 256 bytes
	const MBR_SIZE = 256
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, &mbr); err != nil {
		return fmt.Errorf("Mkdisk: error al serializar el MBR: %v", err)
	}

	data := buf.Bytes()
	if len(data) > MBR_SIZE {
		return fmt.Errorf("Mkdisk: el MBR serializado excede %d bytes", MBR_SIZE)
	}

	// Rellenar con ceros hasta llegar a 512 bytes
	padding := make([]byte, MBR_SIZE-len(data))
	data = append(data, padding...)

	if _, err := file.WriteAt(data, 0); err != nil {
		return fmt.Errorf("Mkdisk: error al escribir el MBR en el disco: %v", err)
	}

	return nil
}

func (d *DiskManager) Rmdisk(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("Rmdisk: el disco en la ruta '%s' no existe", path)
	}
	if filepath.Ext(path) != ".mia" {
		return fmt.Errorf("Rmdisk: el archivo de disco %s no es un archivo .mia", path)
	}
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("Rmdisk: error al eliminar el disco: %v", err)
	}
	return nil
}

func (d *DiskManager) Fdisk(path string, size int64, unit byte, parType byte, fit string, name string) error {

	// 2. Leer el MBR real
	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("Fdisk: error al leer el MBR: %v", err)
	}

	// 2.1 Validar si ya existe una partición con ese nombre en el MBR
	for i := 0; i < 4; i++ {
		existingName := strings.Trim(string(mbr.Mbr_partition[i].Part_name[:]), "\x00")
		if existingName == name {
			return fmt.Errorf("Fdisk: ya existe una partición con el nombre '%s'", name)
		}
	}

	// 2.2 Solo se permite una partición extendida por disco

	var extendedCount int
	for i := 0; i < 4; i++ {
		if mbr.Mbr_partition[i].Part_type == 'E' {
			extendedCount++
		}
	}

	if extendedCount > 1 {
		return fmt.Errorf("Fdisk: solo se permite una partición extendida por disco")
	}

	// 3. Calcular tamaño en bytes
	var partitionSize int64
	switch strings.ToUpper(string(unit)) {
	case "B":
		partitionSize = size
	case "M":
		partitionSize = size * 1024 * 1024
	default: // "K"
		partitionSize = size * 1024
	}

	// 4. Crear partición según tipo
	switch strings.ToUpper(string(parType)) {
	case "P":
		return d.CreatePrimary(name, partitionSize, fit, path)
	case "E":
		return d.CreateExtended(name, partitionSize, fit, path)
	case "L":
		return d.CreateLogical(name, partitionSize, fit, path)
	default:
		return fmt.Errorf("Fdisk: tipo de partición '%c' no reconocido", parType)
	}
}

func (d *DiskManager) Mount(list MountedPartitionList, path string, name string) (MountedPartitionList, error) {

	// --- 1) Verificar si ya está montada (recorriendo montadas) ---
	for _, p := range list.Partitions {
		if p.Path == path && strings.Trim(string(p.Partition.Part_name[:]), "\x00") == name {
			return list, fmt.Errorf("Mount: la partición '%s' en el disco '%s' ya está montada", name, path)
		}
	}

	// --- 2) Abrir y leer el MBR ---
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return list, fmt.Errorf("Mount: no se pudo abrir el disco en '%s'", path)
	}
	defer file.Close()

	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return list, fmt.Errorf("Mount: error al leer el MBR: %v", err)
	}

	// --- 4) Buscar primarias (y rechazar extendidas) sobre el MBR leído ---
	for i := range mbr.Mbr_partition {
		p := &mbr.Mbr_partition[i]
		if strings.Trim(string(p.Part_name[:]), "\x00") == name {
			if p.Part_type == 'E' {
				return list, fmt.Errorf("Mount: no se pueden montar particiones extendidas")
			}

			// 🔹 Generar ID aquí, ya estamos seguros que es montable
			id, partNum := list.GetNextId(path)

			// Actualiza MBR en memoria (montado, correlativo, id)
			p.Part_status = 1
			p.Part_correlative = int64(partNum)

			var pid [4]byte
			copy(pid[:], []byte(id))
			p.Part_id = pid

			if err := d.MbrManager.WriteMBR(&mbr, path); err != nil {
				return list, fmt.Errorf("Mount: error al actualizar el MBR en el disco: %v", err)
			}

			err = list.SetPartition(id, path, *p)
			if err != nil {
				return list, fmt.Errorf("Mount: error al registrar la partición montada: %v", err)
			}

			return list, nil
		}
	}

	// --- 5) Buscar extendida y luego lógicas (EBR) ---
	var extended disk.Partition
	foundExtended := false
	for _, p := range mbr.Mbr_partition {
		if p.Part_type == 'E' {
			extended = p
			foundExtended = true
			break
		}
	}

	if foundExtended {
		currentEBR, err := d.EbrManager.ReadEBR(path, extended.Part_start)
		if err != nil {
			return list, fmt.Errorf("Mount: error al leer el primer EBR: %v", err)
		}
		currentEBRAddress := extended.Part_start

		for {
			ebrName := strings.Trim(string(currentEBR.Part_name[:]), "\x00")
			if ebrName == name {
				currentEBR.Part_mount = 1

				// 🔹 Generar ID aquí
				id, partNum := list.GetNextId(path)

				var pid [4]byte
				copy(pid[:], []byte(id))

				logicalAsPartition := disk.Partition{
					Part_status:      currentEBR.Part_mount,
					Part_type:        'L',
					Part_fit:         currentEBR.Part_fit,
					Part_start:       currentEBR.Part_start,
					Part_size:        currentEBR.Part_size,
					Part_name:        currentEBR.Part_name,
					Part_correlative: int64(partNum),
					Part_id:          pid,
				}

				err = list.SetPartition(id, path, logicalAsPartition)
				if err != nil {
					return list, fmt.Errorf("Mount: error al registrar la partición montada: %v", err)
				}

				if err := d.EbrManager.WriteEBR(&currentEBR, path, currentEBRAddress); err != nil {
					return list, fmt.Errorf("Mount: error al actualizar el EBR en el disco: %v", err)
				}

				return list, nil
			}

			if currentEBR.Part_next == -1 {
				break
			}
			currentEBRAddress = currentEBR.Part_next
			currentEBR, err = d.EbrManager.ReadEBR(path, currentEBRAddress)
			if err != nil {
				return list, fmt.Errorf("Mount: error al leer el EBR: %v", err)
			}
		}
	}

	return list, fmt.Errorf("Mount: error: no se encontró una partición con el nombre '%s' en el disco '%s'", name, path)
}

func (d *DiskManager) Mounted(list MountedPartitionList) (string, error) {
	if len(list.Partitions) == 0 {
		return "", fmt.Errorf("Mounted: no hay particiones montadas")
	}

	var mountedInfo strings.Builder
	header := fmt.Sprintf(
		"| %-10s | %-4s | %-4s | %-8s | %-8s | %-4s | %-6s |\n",
		"Nombre", "ID", "Tipo", "Tamaño", "Inicio", "Fit", "Estado",
	)
	separator := strings.Repeat("-", len(header)) + "\n"
	mountedInfo.WriteString(separator)
	mountedInfo.WriteString(header)
	mountedInfo.WriteString(separator)

	for _, p := range list.Partitions {
		line := fmt.Sprintf(
			"| %-10s | %-4s | %-4c | %-8d | %-8d | %-4c | %-6d |\n",
			strings.Trim(string(p.Partition.Part_name[:]), "\x00"),
			p.Id,
			p.Partition.Part_type,
			p.Partition.Part_size,
			p.Partition.Part_start,
			p.Partition.Part_fit,
			p.Partition.Part_status,
		)
		mountedInfo.WriteString(line)
	}

	mountedInfo.WriteString(separator)
	return mountedInfo.String(), nil
}

func (d *DiskManager) Rep(partition MountedPartition, name string, path string, id string, path_file_ls string, format string) error {
	// Sacar extensión del archivo de salida
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
	if ext == "" {
		return fmt.Errorf("Rep: no se pudo determinar extensión para '%s'", path)
	}

	// Instanciar generador con los managers reales de este FS
	var generator rep.GenerateDiskReport = NewReportDiskManager(
		partition,
	)

	// Mapa de reportes soportados
	reporters := map[string]func(string, string, string) error{
		"mbr":  generator.Mbr,
		"disk": generator.Disk,
	}

	if fn, ok := reporters[name]; ok {
		if err := fn(path, id, ext); err != nil {
			return fmt.Errorf("Rep: error ejecutando reporte '%s': %v", name, err)
		}
		return nil
	}

	return fmt.Errorf("Rep: reporte '%s' no soportado", name)
}
