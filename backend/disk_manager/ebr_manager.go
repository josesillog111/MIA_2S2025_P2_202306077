package dmanager

import (
	disk "backend/disk_structs"
	"encoding/binary"
	"fmt"
	"os"
	"sort"
)

type EbrManager struct {
}

func (e *EbrManager) AddEBR(path string, extended *disk.Partition, newEbr *disk.EBR, ebrs []disk.EBR) error {
	// Abrir el archivo en modo lectura/escritura
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("AddEBR: no se pudo abrir el disco: %v", err)
	}
	defer file.Close()

	// Caso 1: no existen EBRs todavía → el nuevo será el primero
	if len(ebrs) == 0 {
		newEbr.Part_next = -1

		// Escribir el nuevo EBR en el inicio del espacio de la extendida
		if _, err := file.Seek(newEbr.Part_start, 0); err != nil {
			return fmt.Errorf("AddEBR: error al posicionar puntero de archivo: %v", err)
		}
		if err := binary.Write(file, binary.LittleEndian, newEbr); err != nil {
			return fmt.Errorf("AddEBR: error al escribir primer EBR: %v", err)
		}
		return nil
	}

	// Caso 2: ya existen EBRs → insertamos en la lista
	// Ordenar EBRs por Part_start
	sort.Slice(ebrs, func(i, j int) bool {
		return ebrs[i].Part_start < ebrs[j].Part_start
	})

	inserted := false
	for i := range ebrs {
		// si el nuevo EBR va después del actual y antes del siguiente
		if i == len(ebrs)-1 || (newEbr.Part_start > ebrs[i].Part_start && newEbr.Part_start < ebrs[i+1].Part_start) {
			// actualizar el next del actual EBR
			ebrs[i].Part_next = newEbr.Part_start

			// si hay un siguiente, enlazarlo
			if i < len(ebrs)-1 {
				newEbr.Part_next = ebrs[i+1].Part_start
			} else {
				newEbr.Part_next = -1
			}

			// Reescribir el EBR actual con el nuevo Part_next
			if _, err := file.Seek(ebrs[i].Part_start, 0); err != nil {
				return fmt.Errorf("AddEBR: error al posicionar archivo en EBR actual: %v", err)
			}
			if err := binary.Write(file, binary.LittleEndian, &ebrs[i]); err != nil {
				return fmt.Errorf("AddEBR: error al actualizar EBR actual: %v", err)
			}

			// Escribir el nuevo EBR
			if _, err := file.Seek(newEbr.Part_start, 0); err != nil {
				return fmt.Errorf("AddEBR: error al posicionar archivo para nuevo EBR: %v", err)
			}
			if err := binary.Write(file, binary.LittleEndian, newEbr); err != nil {
				return fmt.Errorf("AddEBR: error al escribir nuevo EBR: %v", err)
			}

			inserted = true
			break
		}
	}

	if !inserted {
		return fmt.Errorf("AddEBR: no se pudo insertar el nuevo EBR en la lista")
	}

	return nil
}

func (e *EbrManager) DeleteEBR(path string, ebrToDelete *disk.EBR, ebrs []disk.EBR) error {
	// Abrir el archivo en modo lectura/escritura
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeleteEBR: no se pudo abrir el disco: %v", err)
	}
	defer file.Close()

	// Encontrar el EBR anterior al que queremos eliminar
	var prevEbr *disk.EBR
	for i := range ebrs {
		if ebrs[i].Part_next == ebrToDelete.Part_start {
			prevEbr = &ebrs[i]
			break
		}
	}

	// Si encontramos un EBR anterior, actualizar su Part_next
	if prevEbr != nil {
		prevEbr.Part_next = ebrToDelete.Part_next
	}
	// Reescribir el EBR anterior con el nuevo Part_next
	if prevEbr != nil {
		if _, err := file.Seek(prevEbr.Part_start, 0); err != nil {
			return fmt.Errorf("DeleteEBR: error al posicionar archivo en EBR anterior: %v", err)
		}
		if err := binary.Write(file, binary.LittleEndian, prevEbr); err != nil {
			return fmt.Errorf("DeleteEBR: error al actualizar EBR anterior: %v", err)
		}
	}
	// Si no encontramos un EBR anterior, el EBR a eliminar es el primero
	// En este caso, no necesitamos actualizar ningún otro EBR

	// Opcional: limpiar el EBR eliminado (no es estrictamente necesario)
	var emptyEbr disk.EBR
	if _, err := file.Seek(ebrToDelete.Part_start, 0); err != nil {
		return fmt.Errorf("DeleteEBR: error al posicionar archivo para limpiar EBR: %v", err)
	}
	if err := binary.Write(file, binary.LittleEndian, &emptyEbr); err != nil {
		return fmt.Errorf("DeleteEBR: error al limpiar EBR eliminado: %v", err)
	}
	// Opcional: actualizar la lista de EBRs en memoria
	return nil
}

func (e *EbrManager) ReadEBR(path string, start int64) (disk.EBR, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return disk.EBR{}, fmt.Errorf("ReadEBR: error al abrir el archivo: %v", err)
	}
	defer file.Close()

	_, err = file.Seek(start, 0)
	if err != nil {
		return disk.EBR{}, fmt.Errorf("ReadEBR: error al buscar el inicio del EBR: %v", err)
	}

	var ebr disk.EBR
	err = binary.Read(file, binary.LittleEndian, &ebr)
	if err != nil {
		return disk.EBR{}, fmt.Errorf("ReadEBR: error al leer el EBR: %v", err)
	}
	return ebr, nil
}

func (e *EbrManager) WriteEBR(ebr *disk.EBR, path string, start int64) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("WriteEBR: error al abrir el archivo: %v", err)
	}
	defer file.Close()

	_, err = file.Seek(start, 0)
	if err != nil {
		return fmt.Errorf("WriteEBR: error al buscar el inicio del EBR: %v", err)
	}

	err = binary.Write(file, binary.LittleEndian, ebr)
	if err != nil {
		return fmt.Errorf("WriteEBR: error al escribir el EBR: %v", err)
	}

	return nil
}

func (e *EbrManager) TraverseEBRs(path string, start int64) ([]disk.EBR, error) {
	var ebrs []disk.EBR
	for {
		ebr, err := e.ReadEBR(path, start)
		if err != nil {
			return nil, err
		}
		ebrs = append(ebrs, ebr)
		// Si no hay siguiente, rompemos
		if ebr.Part_next == -1 {
			break
		}
		start = ebr.Part_next
	}
	return ebrs, nil
}

func (e *EbrManager) GetEBRs(path string, extended disk.Partition) ([]disk.EBR, error) {
	var chain []disk.EBR

	if extended.Part_start <= 0 || extended.Part_size <= 0 {
		return chain, nil
	}

	// El primer EBR suele residir exactamente en Part_start de la extendida
	curr := extended.Part_start
	limit := extended.Part_start + extended.Part_size

	for curr > 0 && curr < limit {
		ebr, err := ReadEBRAt(path, curr)
		if err != nil {
			// Si el primer bloque no contiene un EBR válido, salimos silenciosamente
			// o devuelve error si prefieres ser estricto.
			break
		}

		// Evita loops: si todo está vacío, rompe
		if ebr.Part_start <= 0 || ebr.Part_size <= 0 {
			break
		}

		chain = append(chain, *ebr)

		if ebr.Part_next <= 0 {
			break
		}
		curr = ebr.Part_next
	}
	return chain, nil
}

func ReadEBRAt(diskPath string, start int64) (*disk.EBR, error) {
	f, err := os.Open(diskPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := f.Seek(start, 0); err != nil {
		return nil, fmt.Errorf("ReadEBRAt: error al buscar EBR @%d: %w", start, err)
	}

	var ebr disk.EBR
	if err := binary.Read(f, binary.LittleEndian, &ebr); err != nil {
		return nil, fmt.Errorf("ReadEBRAt: error al leer EBR @%d: %w", start, err)
	}
	return &ebr, nil
}
