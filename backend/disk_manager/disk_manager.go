package dmanager

import (
	disk "backend/disk_structs"
	rep "backend/report_manager"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Constantes de control de estado de partición
const (
	STATUS_UNMOUNTED = 0
	STATUS_MOUNTED   = 1
)

// dependencia a partition

type DiskManager struct {
	FreeSpaceManager *FreeSpaceManager
	MbrManager       *MbrManager
	EbrManager       *EbrManager
	PartitionManager *PartitionManager
}

type DiskRecord struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64  `json:"size"`
	Date string `json:"date"`
}

type DiskRegistry struct {
	Disks []DiskRecord `json:"disks"`
}

func NewDiskManager() *DiskManager {
	return &DiskManager{
		FreeSpaceManager: &FreeSpaceManager{},
		MbrManager:       &MbrManager{},
		EbrManager:       &EbrManager{},
		PartitionManager: &PartitionManager{},
	}
}

func (d *DiskManager) diskRegister(nombre, path string, size int64) error {
	regPath := "disks.json"
	var registry DiskRegistry

	// Leer el archivo si existe
	data, err := os.ReadFile(regPath)
	if err == nil {
		if err := json.Unmarshal(data, &registry); err != nil {
			fmt.Printf("diskRegister: advertencia al deserializar JSON: %v. Iniciando con registro vacío.\n", err)
			registry.Disks = []DiskRecord{} // Aseguramos que sea una slice inicializada
		}
	} else if !os.IsNotExist(err) {
		// Error diferente a "archivo no encontrado"
		return fmt.Errorf("diskRegister: error al leer el archivo %s: %v", regPath, err)
	}

	// 1. Buscar si el disco ya existe
	found := false
	for i, disk := range registry.Disks {
		if disk.Name == nombre && disk.Path == path {
			// 2. Si existe, actualizar el registro
			registry.Disks[i].Size = size
			// Actualizar la fecha/hora de registro
			registry.Disks[i].Date = time.Now().Format("2006-01-02 15:04:05")
			found = true
			break
		}
	}

	// 3. Si no existe, agregar un nuevo registro
	if !found {
		registry.Disks = append(registry.Disks, DiskRecord{
			Name: nombre,
			Path: path,
			Size: size,
			Date: time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	// Guardar el archivo actualizado
	data, err = json.MarshalIndent(registry, "", " ")
	if err != nil {
		return fmt.Errorf("diskRegister: error al serializar JSON: %v", err)
	}

	if err := os.WriteFile(regPath, data, 0644); err != nil {
		return fmt.Errorf("diskRegister: error al guardar el archivo: %v", err)
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

	err = d.diskRegister(filepath.Base(path), path, size)
	if err != nil {
		return fmt.Errorf("Mkdisk: error al registrar el disco: %v", err)
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

func (d *DiskManager) Fdisk(path string, size int64, unit byte, parType byte, fit string, delete string, name string, add int64) error {

	// 2. Leer el MBR real
	mbr, err := d.MbrManager.ReadMBR(path)
	if err != nil {
		return fmt.Errorf("Fdisk: error al leer el MBR: %v", err)
	}

	/*
		MODO ADD:
		Este parámetro se utilizará para agregar o quitar
		espacio de la partición. Puede ser positivo o
		negativo. Tomará el parámetro units para las
		unidades a agregar o eliminar.
		En el caso de agregar espacio, deberá comprobar
		que exista espacio libre después de la partición.
		En el caso de quitar espacio se debe comprobar
		que quede espacio en la partición (no espacio
		negativo).
	*/

	if add != 0 {
		// Buscar la partición a modificar
		var targetPartition *disk.Partition

		// Buscar en particiones primarias, extendidas y lógicas
		for i := 0; i < 4; i++ {
			p := &mbr.Mbr_partition[i]
			existingName := strings.Trim(string(p.Part_name[:]), "\x00")
			if p.Part_size > 0 && existingName == name {
				targetPartition = p
				break
			}
		}

		var extendedPartition *disk.Partition // <-- define outside so it's accessible below
		if targetPartition == nil {
			// Buscar en lógicas si no se encontró en primarias/extendidas
			for i := 0; i < 4; i++ {
				p := &mbr.Mbr_partition[i]
				if p.Part_size > 0 && p.Part_type == 'E' {
					extendedPartition = p
					break
				}
			}
			if extendedPartition != nil {
				ebrs, err := d.EbrManager.GetEBRs(path, *extendedPartition)
				if err != nil {
					return fmt.Errorf("Fdisk: error al cargar los EBRs: %v", err)
				}
				for i := range ebrs {
					e := &ebrs[i]
					existingName := strings.Trim(string(e.Part_name[:]), "\x00")

					if existingName == name {
						targetPartition = &disk.Partition{
							Part_status: e.Part_mount,
							Part_type:   'L',
							Part_fit:    e.Part_fit,
							Part_start:  e.Part_start,
							Part_size:   e.Part_size,
						}
						break
					}
				}
			}
		}

		if targetPartition == nil {
			return fmt.Errorf("Fdisk: no se encontró una partición con el nombre '%s' para modificar", name)
		}

		if targetPartition.Part_type == 'P' {
			// Partición primaria
			err := d.PartitionManager.AddPrimary(path, *targetPartition, add, unit)
			if err != nil {
				return fmt.Errorf("Fdisk: error al modificar la partición primaria: %v", err)
			}
		} else if targetPartition.Part_type == 'E' {
			// Partición extendida
			err := d.PartitionManager.AddExtended(path, *targetPartition, add, unit)
			if err != nil {
				return fmt.Errorf("Fdisk: error al modificar la partición extendida: %v", err)
			}
		} else if targetPartition.Part_type == 'L' {
			// Partición lógica
			if extendedPartition == nil {
				// Buscar la extendida si no se encontró antes
				for i := 0; i < 4; i++ {
					p := &mbr.Mbr_partition[i]
					if p.Part_size > 0 && p.Part_type == 'E' {
						extendedPartition = p
						break
					}
				}
				if extendedPartition == nil {
					return fmt.Errorf("Fdisk: no se encontró partición extendida para modificar lógica")
				}
			}
			err := d.PartitionManager.AddLogical(path, *extendedPartition, disk.EBR{
				Part_mount: targetPartition.Part_status,
				Part_fit:   targetPartition.Part_fit,
				Part_start: targetPartition.Part_start,
				Part_size:  targetPartition.Part_size,
				Part_next:  -1,
				Part_name:  [16]byte{},
			}, add, unit, name)
			if err != nil {
				return fmt.Errorf("Fdisk: error al modificar la partición lógica: %v", err)
			}
		} else {
			return fmt.Errorf("Fdisk: el tipo de partición '%c' no es válido para modificación", targetPartition.Part_type)
		}

		return nil

	}

	/*
		MODO DELETE:
		Este parámetro indica que se eliminará una
		partición. Este parámetro se utiliza junto con -name
		y -path.
		Se deberá mostrar un mensaje que permita
		confirmar la eliminación de dicha partición.
		Si la partición no existe deberá mostrar error. Si
		se elimina la partición extendida, deben
		eliminarse las particiones lógicas que tenga
		adentro.
		Recibirá los siguientes valores:
		Fast: Esta opción marca como vacío el espacio en
		la tabla de particiones.
		Full: Esta opción además marcar como vació el
		espacio en la tabla de particiones, rellena el espacio
		con el carácter \0. Si se utiliza otro valor diferente,
		mostrará un mensaje de error.
	*/

	if delete != "" {

		/*

			MENSAJE DE CONFIRMACIÓN  < ACTUALIZAR EN EL FUTURO >

		*/
		fmt.Printf("¿Está seguro que desea eliminar la partición '%s'? (s/n): ", name)
		var resp string
		for {
			fmt.Scanln(&resp)
			resp = strings.ToUpper(resp)
			if resp == "S" {
				break
			}
			if resp == "N" {
				return fmt.Errorf("Fdisk: Eliminación cancelada por el usuario")
			}

			fmt.Print("Respuesta inválida. Por favor ingrese 's' para sí o 'n' para no: ")
		}

		if delete != "FAST" && delete != "FULL" {
			return fmt.Errorf("Fdisk: el modo de borrado '%s' no es válido", delete)
		}

		// Buscar la partición a eliminar
		var targetIndex int = -1
		var targetPartition disk.Partition
		for i := 0; i < 4; i++ {
			p := &mbr.Mbr_partition[i]
			existingName := strings.Trim(string(p.Part_name[:]), "\x00")
			fmt.Println("Comparando con partición:", existingName, " y nombre partición encontrada:", strings.Trim(string(p.Part_name[:]), "\x00"))
			if existingName == name {
				targetIndex = i
				targetPartition = *p
				break
			}
		}
		if targetIndex == -1 {
			return fmt.Errorf("Fdisk: no se encontró una partición con el nombre '%s' para eliminar", name)
		}

		// Eliminar según el tipo de partición

		if mbr.Mbr_partition[targetIndex].Part_type == 'P' {
			// Partición primaria
			err := d.PartitionManager.DeletePrimary(path, targetPartition, delete)
			if err != nil {
				return fmt.Errorf("Fdisk: error al eliminar la partición primaria: %v", err)
			}
		} else if mbr.Mbr_partition[targetIndex].Part_type == 'E' {
			// Partición extendida
			err := d.PartitionManager.DeleteExtended(path, targetPartition, delete)
			if err != nil {
				return fmt.Errorf("Fdisk: error al eliminar la partición extendida: %v", err)
			}
		} else if mbr.Mbr_partition[targetIndex].Part_type == 'L' {
			// Partición lógica
			err := d.PartitionManager.DeleteLogical(path, targetPartition, name, delete)
			if err != nil {
				return fmt.Errorf("Fdisk: error al eliminar la partición lógica: %v", err)
			}
		} else {
			return fmt.Errorf("Fdisk: el tipo de partición '%c' no es válido para eliminación", mbr.Mbr_partition[targetIndex].Part_type)
		}

		return nil
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
		return d.PartitionManager.CreatePrimary(name, partitionSize, fit, path)
	case "E":
		return d.PartitionManager.CreateExtended(name, partitionSize, fit, path)
	case "L":
		return d.PartitionManager.CreateLogical(name, partitionSize, fit, path)
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

			// Generar ID aquí, ya estamos seguros que es montable
			id, partNum := list.GetNextId(path)

			// Actualiza MBR en memoria (montado, correlativo, id)
			p.Part_status = STATUS_MOUNTED
			p.Part_correlative = int64(partNum)

			var pid [4]byte
			copy(pid[:], []byte(id))
			p.Part_id = pid

			// Actualizar MBR en disco
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
			if currentEBR.Part_size > 0 && ebrName == name {
				currentEBR.Part_mount = STATUS_MOUNTED

				// Generar ID aquí
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

func (d *DiskManager) Unmount(list MountedPartitionList, id string) (MountedPartitionList, error) {
	partition := list.GetPartitionById(id)
	if partition == nil {
		return MountedPartitionList{}, fmt.Errorf("Unmount: no existe una partición montada con ID '%s'", id)
	}

	err := list.UnsetPartition(id)
	if err != nil {
		return MountedPartitionList{}, fmt.Errorf("Unmount: error al desmontar la partición: %v", err)
	}

	// Actualizar el MBR en disco para reflejar que la partición ya no está montada
	mbr, err := d.MbrManager.ReadMBR(partition.Path)
	if err != nil {
		return MountedPartitionList{}, fmt.Errorf("Unmount: error al leer el MBR: %v", err)
	}

	for i := 0; i < 4; i++ {
		p := &mbr.Mbr_partition[i]
		if bytes.Equal(p.Part_id[:], []byte(id)) {
			p.Part_status = STATUS_UNMOUNTED
			p.Part_id = [4]byte{}
			p.Part_correlative = 0
			break
		}
	}

	err = d.MbrManager.WriteMBR(&mbr, partition.Path)
	if err != nil {
		return MountedPartitionList{}, fmt.Errorf("Unmount: error al actualizar el MBR en el disco: %v", err)
	}

	return list, nil
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
