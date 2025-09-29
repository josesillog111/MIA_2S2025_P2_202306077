package ext3

import (
	dmanager "backend/disk_manager"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type BlockManager struct {
	Partition dmanager.MountedPartition
	SBManager *SBManager
}

func NewBlockManager(partition dmanager.MountedPartition, sbm *SBManager) *BlockManager {
	return &BlockManager{
		Partition: partition,
		SBManager: sbm,
	}
}

/*

CRUD BLOQUES (YA TIENE INTEGRADO LA MODIFICACIÓN DEL BITMAP)

*/

// Crear bloque
func (bm *BlockManager) CreateBlock(data []byte) (int64, error) {
	blockID, err := bm.AllocateBlock()
	if err != nil {
		return -1, err
	}

	// 2. Escribir los datos en el bloque asignado
	if err := bm.UpdateBlock(blockID, data); err != nil {
		// rollback si falla la escritura
		_ = bm.DeallocateBlock(blockID)
		return -1, err
	}

	return blockID, nil
}

// Leer bloque
func (bm *BlockManager) ReadBlock(id int64) ([]byte, error) {
	blockSize := bm.SBManager.SB.S_block_size
	offset := bm.Partition.Partition.Part_start + bm.SBManager.SB.S_block_start + id*blockSize

	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	data := make([]byte, blockSize)
	if _, err := disk.ReadAt(data, offset); err != nil {
		return nil, fmt.Errorf("ReadBlock: Error leyendo bloque %d: %v", id, err)
	}

	return data, nil
}

// Eliminar bloque
func (bm *BlockManager) DeleteBlock(id int64) error {
	// 1. Opcional: limpiar datos del bloque
	empty := make([]byte, bm.SBManager.SB.S_block_size)
	if err := bm.UpdateBlock(id, empty); err != nil {
		return fmt.Errorf("DeleteBlock: Error limpiando bloque %d: %v", id, err)
	}

	// 2. Liberar en el bitmap
	return bm.DeallocateBlock(id)
}

// Actualizar bloque
func (bm *BlockManager) UpdateBlock(id int64, data []byte) error {
	blockSize := bm.SBManager.SB.S_block_size

	if id < 0 || id >= bm.SBManager.SB.S_blocks_count {
		return fmt.Errorf("UpdateBlock: Id fuera de rango (%d)", id)
	}

	if int64(len(data)) > blockSize {
		return fmt.Errorf("UpdateBlock: Data demasiado grande para el bloque (%d > %d)", len(data), blockSize)
	}

	offset := bm.Partition.Partition.Part_start + bm.SBManager.SB.S_block_start + int64(id)*blockSize

	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("UpdateBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Si data < blockSize, rellenar con ceros
	buf := make([]byte, blockSize)
	copy(buf, data)

	if _, err := disk.WriteAt(buf, offset); err != nil {
		return fmt.Errorf("UpdateBlock: Error escribiendo bloque %d: %v", id, err)
	}

	return nil
}

/*

OPERACIONES DE BAJO NIVEL SOBRE EL BITMAP

*/

// Asignar bloque
func (bm *BlockManager) AllocateBlock() (int64, error) {
	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return -1, fmt.Errorf("AllocateBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Calcula el tamaño del bitmap en bytes
	bitmapSize := (bm.SBManager.SB.S_blocks_count + 7) / 8 // <== revisar este fragmento
	bitmap := make([]byte, bitmapSize)

	// Leer bitmap desde disco (offset absoluto)
	bitmapStart := bm.SBManager.SB.S_bm_block_start
	if _, err := disk.ReadAt(bitmap, bitmapStart); err != nil {
		return -1, fmt.Errorf("AllocateBlock: Error leyendo bitmap bloques: %v", err)
	}

	// Buscar un bloque libre (0 = libre, 1 = ocupado)
	for byteIndex := 0; byteIndex < len(bitmap); byteIndex++ {
		for bit := 0; bit < 8; bit++ {
			blockID := int64(byteIndex*8 + bit)
			if blockID >= bm.SBManager.SB.S_blocks_count {
				break
			}
			if bitmap[byteIndex]&(1<<bit) == 0 {
				// Marcar como ocupado
				bitmap[byteIndex] |= 1 << bit

				// Escribir bitmap actualizado en disco
				if _, err := disk.WriteAt(bitmap, bitmapStart); err != nil {
					return -1, fmt.Errorf("AllocateBlock: Error actualizando bitmap bloques: %v", err)
				}

				// Actualizar superblock
				bm.SBManager.SB.S_free_blocks_count--
				bm.SBManager.WriteSuperBlock(bm.SBManager.SB)
				return blockID, nil
			}
		}
	}

	return -1, fmt.Errorf("AllocateBlock: No hay bloques libres")
}

// Liberar bloque
func (bm *BlockManager) DeallocateBlock(id int64) error {
	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeallocateBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Calcular tamaño del bitmap
	bitmapSize := (bm.SBManager.SB.S_blocks_count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	// Leer bitmap de bloques
	if _, err := disk.ReadAt(bitmap, bm.SBManager.SB.S_bm_block_start); err != nil {
		return fmt.Errorf("DeallocateBlock: Error leyendo bitmap bloques: %v", err)
	}

	// Verificar índices
	if id < 0 || id >= bm.SBManager.SB.S_blocks_count {
		return fmt.Errorf("DeallocateBlock: Id de bloque fuera de rango: %d", id)
	}

	// Liberar bloque (poner en 0)
	byteIndex := id / 8
	bitIndex := id % 8
	bitmap[byteIndex] &^= 1 << bitIndex

	// Escribir bitmap actualizado
	if _, err := disk.WriteAt(bitmap, bm.SBManager.SB.S_bm_block_start); err != nil {
		return fmt.Errorf("DeallocateBlock: Error escribiendo bitmap bloques: %v", err)
	}

	// Actualizar superblock
	bm.SBManager.SB.S_free_blocks_count++
	if err := bm.SBManager.WriteSuperBlock(bm.SBManager.SB); err != nil {
		return fmt.Errorf("DeallocateBlock: Error actualizando superblock: %v", err)
	}

	return nil
}

/*

INICIALIZAR BITMAP

*/

func (bm *BlockManager) IsBlockUsed(id int64) (bool, error) {
	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return false, fmt.Errorf("IsBlockUsed: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Calcular tamaño del bitmap
	bitmapSize := (bm.SBManager.SB.S_blocks_count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	// Leer bitmap de bloques
	if _, err := disk.ReadAt(bitmap, bm.SBManager.SB.S_bm_block_start); err != nil {
		return false, fmt.Errorf("IsBlockUsed: Error leyendo bitmap bloques: %v", err)
	}

	// Verificar índices
	if id < 0 || id >= bm.SBManager.SB.S_blocks_count {
		return false, fmt.Errorf("IsBlockUsed: Id de bloque fuera de rango: %d", id)
	}

	// Comprobar si el bloque está en uso
	byteIndex := id / 8
	bitIndex := id % 8
	bit := (bitmap[byteIndex] >> bitIndex) & 1

	return bit == 1, nil
}

// Inicializar bitmap
func (bm *BlockManager) InitBlockBitmap(offset int64, count int64) error {
	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("InitBlockBitmap: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	realOffset := bm.Partition.Partition.Part_start + offset
	bitmapSize := (count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	// Inicializar en 0 (todos libres, sin reservar manualmente nada)
	for i := range bitmap {
		bitmap[i] = 0
	}

	if _, err := disk.WriteAt(bitmap, realOffset); err != nil {
		return fmt.Errorf("InitBlockBitmap: Error escribiendo bitmap de bloques: %v", err)
	}
	return nil
}

func (bm *BlockManager) GetBitmapBlock(offset int64, count int64) (string, error) {
	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return "", fmt.Errorf("GetBitmapBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	realOffset := bm.SBManager.SB.S_bm_block_start
	bitmapSize := (count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	if _, err := disk.ReadAt(bitmap, realOffset); err != nil {
		return "", fmt.Errorf("GetBitmapBlock: Error leyendo bitmap de bloques: %v", err)
	}

	var builder strings.Builder
	ocupados := 0
	libres := 0

	for i := int64(0); i < count; i++ {
		byteIndex := i / 8
		bitIndex := i % 8
		bit := (bitmap[byteIndex] >> bitIndex) & 1

		// contar
		if bit == 1 {
			ocupados++
		} else {
			libres++
		}

		// construir string (0/1)
		builder.WriteString(fmt.Sprintf("%d\t", bit))

		// separar en filas de 8 bits
		if (i+1)%8 == 0 {
			builder.WriteString(" ")
		}
	}

	result := builder.String()

	return result, nil
}

func (bm *BlockManager) ReadPointers(blockID int64) ([]int64, error) {
	sb := bm.SBManager.SB
	if blockID < 0 || blockID >= sb.S_blocks_count {
		return nil, fmt.Errorf("ReadPointers: id de bloque fuera de rango (%d)", blockID)
	}

	disk, err := os.OpenFile(bm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadPointers: error abriendo disco: %v", err)
	}
	defer disk.Close()

	offset := sb.S_block_start + blockID*sb.S_block_size
	buf := make([]byte, sb.S_block_size)

	if _, err := disk.ReadAt(buf, offset); err != nil {
		return nil, fmt.Errorf("ReadPointers: error leyendo bloque %d: %v", blockID, err)
	}

	// convertir a []int64
	ptrs := make([]int64, sb.S_block_size/8) // 8 bytes por int64
	reader := bytes.NewReader(buf)
	for i := range ptrs {
		if err := binary.Read(reader, binary.LittleEndian, &ptrs[i]); err != nil {
			return nil, fmt.Errorf("ReadPointers: error convirtiendo ptr %d: %v", i, err)
		}
	}

	return ptrs, nil
}
