package dmanager

import (
	disk "backend/disk_structs"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type PartitionManager struct {
	MbrManager       *MbrManager
	EbrManager       *EbrManager
	FreeSpaceManager *FreeSpaceManager
}

func NewPartitionManager() *PartitionManager {
	return &PartitionManager{
		MbrManager:       &MbrManager{},
		EbrManager:       &EbrManager{},
		FreeSpaceManager: &FreeSpaceManager{},
	}
}

/* FUNCIONES PARTICIONES PRIMARIAS */
func (p *PartitionManager) CreatePrimary(name string, size int64, fit string, path string) error {

	// 1. Validaciones
	mbr, err := p.MbrManager.ReadMBR(path)
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
	free := p.FreeSpaceManager.GetFreeSpaces(&mbr)

	var bestFitStart int64 = -1
	switch strings.ToUpper(fit) {
	case "FF":
		bestFitStart = p.FreeSpaceManager.FindFirstFit(size, free)
	case "BF":
		bestFitStart = p.FreeSpaceManager.FindBestFit(size, free)
	case "WF":
		bestFitStart = p.FreeSpaceManager.FindWorstFit(size, free)
	}
	if bestFitStart < 0 {
		return fmt.Errorf("CreatePrimary: no hay espacio suficiente para la partición primaria")
	}

	// 3. Crear la nueva estructura de Partición
	var newPartition disk.Partition
	newPartition.Part_status = STATUS_USED
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
	err = p.MbrManager.WriteMBR(&mbr, path)
	if err != nil {
		return fmt.Errorf("CreatePrimary: error al escribir el MBR: %v", err)
	}

	return nil
}

func (p *PartitionManager) DeletePrimary(path string, partition disk.Partition, delete string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeletePrimary: no se pudo abrir el disco: %v", err)
	}
	defer file.Close()
	mbr, err := p.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("DeletePrimary: error al leer el MBR: %v", err)
	}
	var targetIndex int = -1
	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		existingName := strings.Trim(string(p.Part_name[:]), "\x00")
		if p.Part_status == STATUS_USED && existingName == string(partition.Part_name[:]) {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return fmt.Errorf("DeletePrimary: no se encontró una partición primaria con el nombre '%s'", string(partition.Part_name[:]))
	}
	if delete != "Fast" && delete != "Full" {
		return fmt.Errorf("DeletePrimary: el modo de borrado '%s' no es válido", delete)
	}
	if delete == "Fast" {
		// Marcar como libre en el MBR
		mbr.Mbr_partition[targetIndex] = disk.Partition{}
	}
	if delete == "Full" {
		// Rellenar con ceros en el disco y marcar como libre en el MBR
		start := partition.Part_start
		size := partition.Part_size
		zeroBytes := make([]byte, size)
		if _, err := file.WriteAt(zeroBytes, start); err != nil {
			return fmt.Errorf("DeletePrimary: error al escribir ceros en el disco: %v", err)
		}
		mbr.Mbr_partition[targetIndex] = disk.Partition{}
	}
	// Escribir el MBR actualizado en disco
	err = p.MbrManager.WriteMBR(&mbr, path)
	if err != nil {
		return fmt.Errorf("DeletePrimary: error al escribir el MBR actualizado: %v", err)
	}
	return nil
}

func (p *PartitionManager) AddPrimary(path string, partition disk.Partition, add int64, unit byte) error {
	mbr, err := p.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("AddPrimary: error al leer el MBR: %v", err)
	}
	var targetPartition *disk.Partition
	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		existingName := strings.Trim(string(p.Part_name[:]), "\x00")
		if p.Part_status == STATUS_USED && existingName == string(partition.Part_name[:]) {
			targetPartition = p
			break
		}
	}
	if targetPartition == nil {
		return fmt.Errorf("AddPrimary: no se encontró una partición primaria con el nombre '%s'", string(partition.Part_name[:]))
	}

	// Calcular tamaño a agregar o quitar en bytes
	var delta int64
	switch strings.ToUpper(string(unit)) {
	case "B":
		delta = add
	case "M":
		delta = add * 1024 * 1024
	case "K":
		delta = add * 1024
	default:
		return fmt.Errorf("AddPrimary: unidad '%c' no reconocida", unit)
	}

	// Caso de reducir tamaño
	if delta < 0 {
		// En caso de valor negativo, se debe reducir el tamaño
		if targetPartition.Part_size+delta <= 0 {
			return fmt.Errorf("AddPrimary: no se puede reducir la partición '%s' a tamaño negativo o cero", string(partition.Part_name[:]))
		}
		targetPartition.Part_size += delta
		err = p.MbrManager.WriteMBR(&mbr, path)
		if err != nil {
			return fmt.Errorf("AddPrimary: error al escribir el MBR actualizado: %v", err)
		}
		return nil
	}
	// Caso de aumentar tamaño
	if delta > 0 {
		// Verificar espacio libre después de la partición
		endOfPartition := targetPartition.Part_start + targetPartition.Part_size
		freeSpaces := p.FreeSpaceManager.GetFreeSpaces(&mbr)
		hasSpace := false
		for _, fs := range freeSpaces {
			if fs.Start >= endOfPartition && fs.Size >= delta {
				hasSpace = true
				break
			}
		}
		if !hasSpace {
			return fmt.Errorf("AddPrimary: no hay espacio libre suficiente después de la partición '%s' para agregar %d bytes", string(partition.Part_name[:]), delta)
		}
		// Modificar el tamaño de la partición
		targetPartition.Part_size += delta
		// Escribir el MBR actualizado en disco
		err = p.MbrManager.WriteMBR(&mbr, path)
		if err != nil {
			return fmt.Errorf("AddPrimary: error al escribir el MBR actualizado: %v", err)
		}
		return nil
	}
	return nil
}

/* FUNCIONES PARTICIONES EXTENDIDAS */
func (p *PartitionManager) CreateExtended(name string, size int64, fit string, path string) error {
	mbr, err := p.MbrManager.ReadMBR(path)
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
	freeSpaces := p.FreeSpaceManager.GetFreeSpaces(&mbr)
	var bestFitStart int64 = -1

	switch strings.ToUpper(fit) {
	case "FF":
		bestFitStart = p.FreeSpaceManager.FindFirstFit(size, freeSpaces)
	case "BF":
		bestFitStart = p.FreeSpaceManager.FindBestFit(size, freeSpaces)
	case "WF":
		bestFitStart = p.FreeSpaceManager.FindWorstFit(size, freeSpaces)
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
	err = p.MbrManager.WriteMBR(&mbr, path)
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

	err = p.EbrManager.WriteEBR(&firstEBR, path, newPartition.Part_start)
	if err != nil {
		return fmt.Errorf("CreateExtended: error al inicializar el primer EBR: %v", err)
	}

	return nil
}

func (p *PartitionManager) DeleteExtended(path string, extended disk.Partition, delete string) error {
	// Eliminar todas las lógicas dentro de la extendida
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeleteExtended: no se pudo abrir el disco: %v", err)
	}
	defer file.Close()
	ebrs, err := p.EbrManager.GetEBRs(path, extended)
	if err != nil {
		return fmt.Errorf("DeleteExtended: error al cargar los EBRs: %v", err)
	}
	for _, e := range ebrs {
		existingName := strings.Trim(string(e.Part_name[:]), "\x00")
		if e.Part_mount == STATUS_USED {
			err := p.DeleteLogical(path, extended, existingName, delete)
			if err != nil {
				return fmt.Errorf("DeleteExtended: error al eliminar la lógica '%s': %v", existingName, err)
			}
		}
	}
	mbr, err := p.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("DeleteExtended: error al leer el MBR: %v", err)
	}
	var targetIndex int = -1
	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		existingName := strings.Trim(string(p.Part_name[:]), "\x00")
		if p.Part_status == STATUS_USED && existingName == string(extended.Part_name[:]) {
			targetIndex = i
			break
		}
	}
	if targetIndex == -1 {
		return fmt.Errorf("DeleteExtended: no se encontró una partición extendida con el nombre '%s'", string(extended.Part_name[:]))
	}
	if delete != "Fast" && delete != "Full" {
		return fmt.Errorf("DeleteExtended: el modo de borrado '%s' no es válido", delete)
	}
	if delete == "Fast" {
		// Marcar como libre en el MBR
		mbr.Mbr_partition[targetIndex] = disk.Partition{}
	}
	if delete == "Full" {
		// Rellenar con ceros en el disco y marcar como libre en el MBR
		start := extended.Part_start
		size := extended.Part_size
		zeroBytes := make([]byte, size)
		if _, err := file.WriteAt(zeroBytes, start); err != nil {
			return fmt.Errorf("DeleteExtended: error al escribir ceros en el disco: %v", err)
		}
		mbr.Mbr_partition[targetIndex] = disk.Partition{}
	}
	// Escribir el MBR actualizado en disco
	err = p.MbrManager.WriteMBR(&mbr, path)
	if err != nil {
		return fmt.Errorf("DeleteExtended: error al escribir el MBR: %v", err)
	}
	return nil
}

func (p *PartitionManager) AddExtended(path string, partition disk.Partition, add int64, unit byte) error {
	mbr, err := p.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("AddExtended: error al leer el MBR: %v", err)
	}
	var targetPartition *disk.Partition
	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		existingName := strings.Trim(string(p.Part_name[:]), "\x00")
		if p.Part_status == STATUS_USED && existingName == string(partition.Part_name[:]) {
			targetPartition = p
			break
		}
	}
	if targetPartition == nil {
		return fmt.Errorf("AddExtended: no se encontró una partición extendida con el nombre '%s'", string(partition.Part_name[:]))
	}
	// Calcular tamaño a agregar o quitar en bytes
	var delta int64
	switch strings.ToUpper(string(unit)) {
	case "B":
		delta = add
	case "M":
		delta = add * 1024 * 1024
	case "K":
		delta = add * 1024
	default:
		return fmt.Errorf("AddExtended: unidad '%c' no reconocida", unit)
	}
	// Caso de reducir tamaño
	if delta < 0 {
		// En caso de valor negativo, se debe reducir el tamaño
		if targetPartition.Part_size+delta <= 0 {
			return fmt.Errorf("AddExtended: no se puede reducir la partición '%s' a tamaño negativo o cero", string(partition.Part_name[:]))
		}
		targetPartition.Part_size += delta
		err = p.MbrManager.WriteMBR(&mbr, path)
		if err != nil {
			return fmt.Errorf("AddExtended: error al escribir el MBR actualizado: %v", err)
		}
		return nil
	}
	// Caso de aumentar tamaño
	if delta > 0 {
		// Verificar espacio libre después de la partición
		endOfPartition := targetPartition.Part_start + targetPartition.Part_size
		freeSpaces := p.FreeSpaceManager.GetFreeSpaces(&mbr)
		hasSpace := false
		for _, fs := range freeSpaces {
			if fs.Start >= endOfPartition && fs.Size >= delta {
				hasSpace = true
				break
			}
		}
		if !hasSpace {
			return fmt.Errorf("AddExtended: no hay espacio libre suficiente después de la partición '%s' para agregar %d bytes", string(partition.Part_name[:]), delta)
		}
		// Modificar el tamaño de la partición
		targetPartition.Part_size += delta
		// Escribir el MBR actualizado en disco
		err = p.MbrManager.WriteMBR(&mbr, path)
		if err != nil {
			return fmt.Errorf("AddExtended: error al escribir el MBR actualizado: %v", err)
		}
		return nil
	}
	return nil
}

/* FUNCIONES PARTICIONES EXTENDIDAS */
func (p *PartitionManager) CreateLogical(name string, size int64, fit string, path string) error {
	// 1. Buscar la partición extendida en el MBR
	mbr, err := p.MbrManager.ReadMBR(path)
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
	ebrs, err := p.EbrManager.GetEBRs(path, *extendedPartition)
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
	free := p.FreeSpaceManager.GetFreeSpacesInExtended(*extendedPartition, ebrs)

	var start int64 = -1
	switch strings.ToUpper(fit) {
	case "FF":
		start = p.FreeSpaceManager.FindFirstFit(size, free)
	case "BF":
		start = p.FreeSpaceManager.FindBestFit(size, free)
	case "WF":
		start = p.FreeSpaceManager.FindWorstFit(size, free)
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
	err = p.EbrManager.AddEBR(path, extendedPartition, &newEbr, ebrs)
	if err != nil {
		return fmt.Errorf("CreateLogical: error al insertar el nuevo EBR: %v", err)
	}

	return nil
}

func (p *PartitionManager) DeleteLogical(path string, extended disk.Partition, name string, delete string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeleteLogical: no se pudo abrir el disco: %v", err)
	}
	defer file.Close()

	ebrs, err := p.EbrManager.GetEBRs(path, extended)
	if err != nil {
		return fmt.Errorf("DeleteLogical: error al cargar los EBRs: %v", err)
	}

	var targetEBR *disk.EBR
	var previousEBR *disk.EBR

	for i := range ebrs {
		e := &ebrs[i]
		existingName := strings.Trim(string(e.Part_name[:]), "\x00")
		if e.Part_mount == STATUS_USED && existingName == name {
			targetEBR = e
			if i > 0 {
				previousEBR = &ebrs[i-1]
			}
			break
		}
	}

	if targetEBR == nil {
		return fmt.Errorf("DeleteLogical: no se encontró una partición lógica con el nombre '%s'", name)
	}

	if delete != "Fast" && delete != "Full" {
		return fmt.Errorf("DeleteLogical: el modo de borrado '%s' no es válido", delete)
	}

	if delete == "Fast" {
		targetEBR.Part_mount = STATUS_FREE
	} else if delete == "Full" {
		// Rellenar con ceros
		start := targetEBR.Part_start
		size := targetEBR.Part_size
		zeroBytes := make([]byte, size)
		if _, err := file.WriteAt(zeroBytes, start); err != nil {
			return fmt.Errorf("DeleteLogical: error al escribir ceros en el disco: %v", err)
		}
		targetEBR.Part_mount = STATUS_FREE
	}

	// Reajustar el enlace con el anterior
	if previousEBR != nil {
		previousEBR.Part_next = targetEBR.Part_next
		if _, err := file.Seek(previousEBR.Part_start, 0); err != nil {
			return fmt.Errorf("DeleteLogical: error al posicionar archivo en EBR anterior: %v", err)
		}
		if err := binary.Write(file, binary.LittleEndian, previousEBR); err != nil {
			return fmt.Errorf("DeleteLogical: error al actualizar EBR anterior: %v", err)
		}
	}

	// Reescribir el EBR actual
	if _, err := file.Seek(targetEBR.Part_start, 0); err != nil {
		return fmt.Errorf("DeleteLogical: error al posicionar archivo en EBR actual: %v", err)
	}
	if err := binary.Write(file, binary.LittleEndian, targetEBR); err != nil {
		return fmt.Errorf("DeleteLogical: error al actualizar EBR actual: %v", err)
	}

	return nil
}

func (p *PartitionManager) AddLogical(path string, extended disk.Partition, partition disk.EBR, add int64, unit byte, name string) error {
	ebrs, err := p.EbrManager.GetEBRs(path, extended)
	if err != nil {
		return fmt.Errorf("AddLogical: error al cargar los EBRs: %v", err)
	}

	var targetEBR *disk.EBR
	for i := range ebrs {
		e := &ebrs[i]
		existingName := strings.Trim(string(e.Part_name[:]), "\x00")
		if e.Part_mount == STATUS_USED && existingName == name {
			targetEBR = e
			break
		}
	}
	if targetEBR == nil {
		return fmt.Errorf("AddLogical: no se encontró una partición lógica con el nombre '%s'", name)
	}

	// Calcular delta
	var delta int64
	switch strings.ToUpper(string(unit)) {
	case "B":
		delta = add
	case "M":
		delta = add * 1024 * 1024
	case "K":
		delta = add * 1024
	default:
		return fmt.Errorf("AddLogical: unidad '%c' no reconocida", unit)
	}

	if delta < 0 {
		if targetEBR.Part_size+delta <= 0 {
			return fmt.Errorf("AddLogical: no se puede reducir la partición '%s' a tamaño negativo o cero", name)
		}
		targetEBR.Part_size += delta
		return p.EbrManager.WriteEBR(targetEBR, path, targetEBR.Part_start)
	}

	if delta > 0 {
		endOfPartition := targetEBR.Part_start + targetEBR.Part_size
		freeSpaces := p.FreeSpaceManager.GetFreeSpacesInExtended(extended, ebrs)
		hasSpace := false
		for _, fs := range freeSpaces {
			if fs.Start >= endOfPartition && fs.Size >= delta {
				hasSpace = true
				break
			}
		}
		if !hasSpace {
			return fmt.Errorf("AddLogical: no hay espacio libre suficiente después de la partición lógica '%s'", name)
		}
		targetEBR.Part_size += delta
		return p.EbrManager.WriteEBR(targetEBR, path, targetEBR.Part_start)
	}

	return nil
}
