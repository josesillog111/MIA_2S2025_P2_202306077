package dmanager

import (
	disk "backend/disk_structs"
	fsys "backend/file_system"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	partitionNumbers      = make(map[string]int)  // guarda el correlativo de particiones por disco
	diskLetters           = make(map[string]byte) // guarda la letra asignada a cada disco
	nextLetter       byte = 'A'
)

type MountedPartition struct {
	Id         string
	Path       string
	FileSystem fsys.FileSystem
	Partition  *disk.Partition
}

func (mp *MountedPartition) HasFileSystem() (bool, error) {
	if mp.Partition == nil {
		return false, fmt.Errorf("HasFileSystem: partición no inicializada")
	}

	disk, err := os.OpenFile(mp.Path, os.O_RDONLY, 0666)
	if err != nil {
		return false, fmt.Errorf("HasFileSystem: error abriendo disco: %v", err)
	}
	defer disk.Close()

	// El superbloque suele estar al inicio de la partición (ajustar según tu implementación)
	offset := mp.Partition.Part_start
	buf := make([]byte, 64) // suficiente para leer el campo identificador
	if _, err := disk.ReadAt(buf, offset); err != nil {
		return false, fmt.Errorf("HasFileSystem: error leyendo superbloque: %v", err)
	}

	// Leer firma en los primeros bytes
	signature := string(buf[:4]) // "EXT2"
	if signature == "EXT2" {
		return true, nil
	}

	return false, nil
}

type MountedPartitionList struct {
	Partitions []*MountedPartition
}

func NewMountedPartitionList() *MountedPartitionList {
	return &MountedPartitionList{
		Partitions: make([]*MountedPartition, 0),
	}
}

func (m *MountedPartitionList) GetPartitionByPathAndName(path string, name string) *MountedPartition {
	for _, part := range m.Partitions {
		if part.Path == path && part.Partition != nil && strings.TrimRight(string(part.Partition.Part_name[:]), "\x00") == name {
			return part
		}
	}
	return nil
}

func (m *MountedPartitionList) GetNextId(path string) (string, int64) {
	carnet := "202306077"
	lastTwo := carnet[len(carnet)-2:]

	// Verificar si ya existe letra asignada para este disco
	letter, exists := diskLetters[path]
	if !exists {
		// Si es disco nuevo, asignar siguiente letra disponible
		letter = nextLetter
		nextLetter++
		diskLetters[path] = letter
		partitionNumbers[path] = 1
	} else {
		// Si el disco ya existe, incrementar número de partición
		partitionNumbers[path]++
	}

	// Obtener número actual de partición
	partNum := partitionNumbers[path]

	id := lastTwo + strconv.Itoa(partNum) + strings.ToUpper(string(letter))

	return id, int64(partNum)
}

func (m *MountedPartitionList) SetPartition(id string, path string, part disk.Partition) error {
	m.Partitions = append(m.Partitions, &MountedPartition{
		Id:        id,
		Path:      path,
		Partition: &part,
	})
	return nil
}

func (m *MountedPartitionList) UnsetPartition(id string) error {
	for i, part := range m.Partitions {
		if part.Id == id {
			// Eliminar la partición montada de la lista
			m.Partitions = append(m.Partitions[:i], m.Partitions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("UnsetPartition: no se encontró la partición con ID %s", id)
}

func (m *MountedPartitionList) GetPartitionById(id string) *MountedPartition {
	for _, part := range m.Partitions {
		if part.Id == id {
			return part
		}
	}
	return nil
}

func (m *MountedPartitionList) GetPathById(id string) string {
	for _, part := range m.Partitions {
		if part.Id == id {
			return part.Path
		}
	}
	return ""
}

func (m *MountedPartitionList) GetPartitions() string {
	var result strings.Builder
	result.WriteString("Particiones montadas:\n")
	for _, part := range m.Partitions {
		partName := ""
		if part.Partition != nil {
			partName = strings.TrimRight(string(part.Partition.Part_name[:]), "\x00")
		}
		result.WriteString(fmt.Sprintf("ID: %s, Disco: %s, Partición: %s\n", part.Id, part.Path, partName))
	}
	return result.String()
}
