package dmanager

import (
	disk "backend/disk_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type MbrManager struct {
}

// read
const MBR_SIZE = 256

func (m *MbrManager) ReadMBR(path string) (disk.MBR, error) {
	var mbr disk.MBR
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return mbr, fmt.Errorf("ReadMBR: error al abrir el disco: %v", err)
	}
	defer file.Close()

	// Leer los primeros 512 bytes
	data := make([]byte, MBR_SIZE)
	if _, err := file.ReadAt(data, 0); err != nil {
		return mbr, fmt.Errorf("ReadMBR: error al leer el MBR: %v", err)
	}

	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &mbr); err != nil {
		return mbr, fmt.Errorf("ReadMBR: error al deserializar el MBR: %v", err)
	}

	return mbr, nil
}

func (m *MbrManager) WriteMBR(mbr *disk.MBR, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("WriteMBR: error al abrir el disco para escribir el MBR: %v", err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, mbr); err != nil {
		return fmt.Errorf("WriteMBR: error al serializar el MBR: %v", err)
	}

	data := buf.Bytes()
	if len(data) > MBR_SIZE {
		return fmt.Errorf("WriteMBR: el MBR serializado excede %d bytes", MBR_SIZE)
	}

	// Rellenar hasta 512 bytes
	padding := make([]byte, MBR_SIZE-len(data))
	data = append(data, padding...)

	if _, err := file.WriteAt(data, 0); err != nil {
		return fmt.Errorf("WriteMBR: error al escribir el MBR en el disco: %v", err)
	}

	return nil
}

func (m *MbrManager) GetPartitionOnMBR(path string) (*[]disk.Partition, error) {
	MBR, err := m.ReadMBR(path)
	if err != nil {
		return nil, fmt.Errorf("WriteMBR: error al leer el MBR: %v", err)
	}
	partitions := MBR.Mbr_partition[:]
	return &partitions, nil
}
