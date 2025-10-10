package ext3

import (
	dmanager "backend/disk_manager"
	file "backend/ext3_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

type FileManager struct {
	Partition        dmanager.MountedPartition
	BlockManager     *BlockManager
	InodeManager     *InodeManager
	DirectoryManager *DirectoryManager
	SBManager        *SBManager
}

func NewFileManager(partition dmanager.MountedPartition, sbm *SBManager) *FileManager {
	bm := NewBlockManager(partition, sbm)
	im := NewInodeManager(partition, sbm)
	dm := NewDirectoryManager(partition, sbm)

	return &FileManager{
		Partition:        partition,
		BlockManager:     bm,
		InodeManager:     im,
		DirectoryManager: dm,
		SBManager:        sbm,
	}
}

/*
	Funciones auxiliares
*/

// Dividir un slice de bytes en chunks de tamaño fijo
func (fm *FileManager) chunkBytes(data []byte, size int) [][]byte {
	var chunks [][]byte
	for start := 0; start < len(data); start += size {
		end := start + size
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[start:end])
	}
	// Si data estaba vacío, devuelve 0 chunks
	return chunks
}

// Convierte un struct Pointer a un slice de bytes
func (fm *FileManager) pointerToBytes(p file.Pointer) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Leer bloques indirectos recursivamente
func (fm *FileManager) readIndirect(blockID int64, level int, out *[]byte) error {
	if blockID == -1 {
		return nil
	}

	// Leer bloque de punteros
	raw, err := fm.BlockManager.ReadBlock(blockID)
	if err != nil {
		return fmt.Errorf("ReadIndirect: Error leyendo bloque %d: %v", blockID, err)
	}

	var ptr file.Pointer
	buf := bytes.NewReader(raw)
	if err := binary.Read(buf, binary.LittleEndian, &ptr); err != nil {
		return fmt.Errorf("ReadIndirect: Error parseando bloque de punteros: %v", err)
	}

	// Recorrer los punteros
	for _, p := range ptr.B_pointers {
		if p == -1 {
			continue
		}
		if level == 1 {
			// Indirecto simple → apunta a bloques de datos
			data, err := fm.BlockManager.ReadBlock(p)
			if err != nil {
				return fmt.Errorf("ReadIndirect: Error leyendo bloque de datos %d: %v", p, err)
			}
			*out = append(*out, data...)
		} else {
			// Indirecto doble o triple → recursivo
			if err := fm.readIndirect(p, level-1, out); err != nil {
				return err
			}
		}
	}

	return nil
}

// Escribir bloques indirectos recursivamente
func (fm *FileManager) writeIndirect(inode *file.Inode, index int, level int, chunks [][]byte, chunkIndex *int) error {
	blockID := inode.I_block[index]
	var ptr file.Pointer

	if blockID == -1 {
		// crear bloque de punteros vacío
		emptyPtr := file.Pointer{}
		for i := range emptyPtr.B_pointers {
			emptyPtr.B_pointers[i] = -1
		}
		raw := new(bytes.Buffer)
		if err := binary.Write(raw, binary.LittleEndian, &emptyPtr); err != nil {
			return err
		}
		newBlock, err := fm.BlockManager.CreateBlock(raw.Bytes())
		if err != nil {
			return err
		}
		inode.I_block[index] = newBlock
		blockID = newBlock
	}

	// leer bloque de punteros
	raw, err := fm.BlockManager.ReadBlock(blockID)
	if err != nil {
		return err
	}
	buf := bytes.NewReader(raw)
	if err := binary.Read(buf, binary.LittleEndian, &ptr); err != nil {
		return err
	}

	// recorrer punteros
	for i := 0; i < len(ptr.B_pointers) && *chunkIndex < len(chunks); i++ {
		if level == 1 {
			// indirecto simple → bloques de datos
			if ptr.B_pointers[i] == -1 {
				newBlock, err := fm.BlockManager.CreateBlock(chunks[*chunkIndex])
				if err != nil {
					return err
				}
				ptr.B_pointers[i] = newBlock
			} else {
				if err := fm.BlockManager.UpdateBlock(ptr.B_pointers[i], chunks[*chunkIndex]); err != nil {
					return err
				}
			}
			*chunkIndex++
		} else {
			// doble o triple → recursión
			if ptr.B_pointers[i] == -1 {
				// crear nuevo bloque de punteros vacío
				emptyPtr := file.Pointer{}
				for j := range emptyPtr.B_pointers {
					emptyPtr.B_pointers[j] = -1
				}
				raw := new(bytes.Buffer)
				if err := binary.Write(raw, binary.LittleEndian, &emptyPtr); err != nil {
					return err
				}
				newBlock, err := fm.BlockManager.CreateBlock(raw.Bytes())
				if err != nil {
					return err
				}
				ptr.B_pointers[i] = newBlock
			}
			if err := fm.writeIndirect(inode, -1, level-1, chunks, chunkIndex); err != nil {
				return err
			}
		}
	}

	// guardar bloque de punteros actualizado
	bufOut := new(bytes.Buffer)
	if err := binary.Write(bufOut, binary.LittleEndian, &ptr); err != nil {
		return err
	}
	if err := fm.BlockManager.UpdateBlock(blockID, bufOut.Bytes()); err != nil {
		return err
	}

	return nil
}

/*
	CRUD MANEJO DE LAS ESTRUCTURAS
	"ARCHIVOS". TAMBIÉN SE MODIFICAN LOS BITMAPS
*/

// Crear archivo
func (fm *FileManager) CreateFile(parentInodeID int64, name string, content []byte, owner string, perms uint16) (int64, error) {
	blockSize := int(fm.SBManager.SB.S_block_size)

	// fragmentar contenido
	chunks := fm.chunkBytes(content, blockSize)
	// 1) Crear inodo
	inodeID, inode, err := fm.InodeManager.CreateInode(FILE_TYPE, int64(len(content)), owner, perms)
	if err != nil {
		return -1, fmt.Errorf("CreateFile: Error creando inodo: %v", err)
	}

	// track para rollback
	var createdBlocks []int64
	cleanup := func() {
		for i := len(createdBlocks) - 1; i >= 0; i-- {
			_ = fm.BlockManager.DeleteBlock(createdBlocks[i])
		}
		_ = fm.InodeManager.DeleteInode(inodeID)
	}

	// 2) Directos
	chunkIdx := 0
	for i := 0; i < 12 && chunkIdx < len(chunks); i++ {
		blockID, err := fm.BlockManager.CreateBlock(chunks[chunkIdx])
		if err != nil {
			cleanup()
			return -1, fmt.Errorf("CreateFile: Error creando bloque directo: %v", err)
		}
		createdBlocks = append(createdBlocks, blockID)
		inode.I_block[i] = blockID
		chunkIdx++
	}

	// 3) Indirecto simple
	if chunkIdx < len(chunks) {
		var p file.Pointer
		for i := range p.B_pointers {
			p.B_pointers[i] = -1
		}
		for i := 0; i < 16 && chunkIdx < len(chunks); i++ {
			blockID, err := fm.BlockManager.CreateBlock(chunks[chunkIdx])
			if err != nil {
				cleanup()
				return -1, fmt.Errorf("CreateFile: Error creando bloque (indirecto simple - datos): %v", err)
			}
			createdBlocks = append(createdBlocks, blockID)
			p.B_pointers[i] = blockID
			chunkIdx++
		}
		pb, err := fm.pointerToBytes(p)
		if err != nil {
			cleanup()
			return -1, fmt.Errorf("CreateFile: Error serializando pointer simple: %v", err)
		}
		ptrBlockID, err := fm.BlockManager.CreateBlock(pb)
		if err != nil {
			cleanup()
			return -1, fmt.Errorf("CreateFile: Error creando bloque pointer simple: %v", err)
		}
		createdBlocks = append(createdBlocks, ptrBlockID)
		inode.I_block[12] = ptrBlockID
	}

	// 4) Indirecto doble
	if chunkIdx < len(chunks) {
		var lvl1 file.Pointer
		for i := range lvl1.B_pointers {
			lvl1.B_pointers[i] = -1
		}

		for i := 0; i < 16 && chunkIdx < len(chunks); i++ {
			var lvl2 file.Pointer
			for j := range lvl2.B_pointers {
				lvl2.B_pointers[j] = -1
			}

			for j := 0; j < 16 && chunkIdx < len(chunks); j++ {
				blockID, err := fm.BlockManager.CreateBlock(chunks[chunkIdx])
				if err != nil {
					cleanup()
					return -1, fmt.Errorf("CreateFile: Error creando bloque (doble - datos): %v", err)
				}
				createdBlocks = append(createdBlocks, blockID)
				lvl2.B_pointers[j] = blockID
				chunkIdx++
			}

			// persistir lvl2
			lvl2Bytes, err := fm.pointerToBytes(lvl2)
			if err != nil {
				cleanup()
				return -1, fmt.Errorf("CreateFile: Error serializando lvl2: %v", err)
			}
			lvl2ID, err := fm.BlockManager.CreateBlock(lvl2Bytes)
			if err != nil {
				cleanup()
				return -1, fmt.Errorf("CreateFile: Error creando bloque lvl2 pointer: %v", err)
			}
			createdBlocks = append(createdBlocks, lvl2ID)
			lvl1.B_pointers[i] = lvl2ID
		}

		// persistir lvl1
		lvl1Bytes, err := fm.pointerToBytes(lvl1)
		if err != nil {
			cleanup()
			return -1, fmt.Errorf("CreateFile: Error serializando lvl1: %v", err)
		}
		lvl1ID, err := fm.BlockManager.CreateBlock(lvl1Bytes)
		if err != nil {
			cleanup()
			return -1, fmt.Errorf("CreateFile: Error creando bloque lvl1 pointer: %v", err)
		}
		createdBlocks = append(createdBlocks, lvl1ID)
		inode.I_block[13] = lvl1ID
	}

	// 5) Persistir inodo
	if err := fm.InodeManager.UpdateInode(inodeID, inode); err != nil {
		cleanup()
		return -1, fmt.Errorf("CreateFile: Error actualizando inodo final: %v", err)
	}

	// 6) Añadir entrada en directorio padre
	if err := fm.DirectoryManager.AddEntry(parentInodeID, name, inodeID); err != nil {
		cleanup()
		return -1, fmt.Errorf("CreateFile: Error agregando entrada al padre: %v", err)
	}

	return inodeID, nil
}

// Leer contenido del archivo
func (fm *FileManager) ReadFile(inodeID int64) ([]byte, error) {

	inode, err := fm.InodeManager.ReadInode(inodeID)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: Error leyendo inodo %d: %v", inodeID, err)
	}

	if byte(inode.I_type) != FILE_TYPE {
		return nil, fmt.Errorf("ReadFile: El inodo %d no es un archivo", inodeID)
	}

	var fileData []byte

	// 🔹 1. Leer bloques directos
	for _, blockID := range inode.I_block[:12] {
		if blockID != -1 {
			data, err := fm.BlockManager.ReadBlock(blockID)
			if err != nil {
				return nil, fmt.Errorf("ReadFile: Error leyendo bloque directo %d: %v", blockID, err)
			}
			fileData = append(fileData, data...)
		}
	}

	// 🔹 2. Leer bloques indirectos simples
	if inode.I_block[12] != -1 {
		if err := fm.readIndirect(inode.I_block[12], 1, &fileData); err != nil {
			return nil, fmt.Errorf("ReadFile: Error en indirecto simple: %v", err)
		}
	}

	// 🔹 3. Leer bloques indirectos dobles
	if inode.I_block[13] != -1 {
		if err := fm.readIndirect(inode.I_block[13], 2, &fileData); err != nil {
			return nil, fmt.Errorf("ReadFile: Error en indirecto doble: %v", err)
		}
	}

	// 🔹 4. Leer bloques indirectos triples
	if inode.I_block[14] != -1 {
		if err := fm.readIndirect(inode.I_block[14], 3, &fileData); err != nil {
			return nil, fmt.Errorf("ReadFile: Error en indirecto triple: %v", err)
		}
	}

	// 🔹 5. Ajustar al tamaño real
	if int64(len(fileData)) > inode.I_size {
		fileData = fileData[:inode.I_size]
	}

	return fileData, nil
}

// Eliminar archivo
func (fm *FileManager) DeleteFile(parentInodeID int64, name string) error {
	// 1. Buscar la entrada en el directorio padre
	entries, err := fm.DirectoryManager.ReadDirectory(parentInodeID)
	if err != nil {
		return fmt.Errorf("DeleteFile: Error leyendo directorio padre: %v", err)
	}

	var inodeID int64 = -1
	for _, entry := range entries {
		entryName := strings.Trim(string(entry.B_name[:]), "\x00")
		if entryName == name {
			inodeID = entry.B_inodo
			break
		}
	}
	if inodeID == -1 {
		return fmt.Errorf("DeleteFile: El archivo '%s' no existe en el directorio", name)
	}

	// 2. Leer el inodo del archivo
	inode, err := fm.InodeManager.ReadInode(inodeID)
	if err != nil {
		return fmt.Errorf("DeleteFile: Error leyendo inodo del archivo: %v", err)
	}
	if byte(inode.I_type) != FILE_TYPE {
		return fmt.Errorf("DeleteFile: El inodo %d no es un archivo", inodeID)
	}

	// 3. Liberar bloques de contenido
	for _, blockID := range inode.I_block {
		if blockID != -1 {
			if err := fm.BlockManager.DeallocateBlock(blockID); err != nil {
				return fmt.Errorf("DeleteFile: Error liberando bloque %d: %v", blockID, err)
			}
		}
	}

	// 4. Eliminar entrada del directorio padre
	if err := fm.DirectoryManager.DeleteEntry(parentInodeID, name); err != nil {
		return fmt.Errorf("DeleteFile: Error eliminando entrada del directorio padre: %v", err)
	}

	// 5. Eliminar inodo
	if err := fm.InodeManager.DeleteInode(inodeID); err != nil {
		return fmt.Errorf("DeleteFile: Error eliminando inodo: %v", err)
	}

	return nil
}

// Actualizar contenido del archivo
func (fm *FileManager) UpdateFile(inodeID int64, newContent []byte) error {
	inode, err := fm.InodeManager.ReadInode(inodeID)
	if err != nil {
		return fmt.Errorf("UpdateFile: Error leyendo inodo: %v", err)
	}
	if inode.I_type != FILE_TYPE {
		return fmt.Errorf("UpdateFile: El inodo %d no es un archivo", inodeID)
	}

	blockSize := int(fm.SBManager.SB.S_block_size)
	var chunks [][]byte
	for len(newContent) > 0 {
		end := blockSize
		if len(newContent) < end {
			end = len(newContent)
		}
		chunks = append(chunks, newContent[:end])
		newContent = newContent[end:]
	}

	// Escribir en bloques directos
	chunkIndex := 0
	for i := 0; i < 12 && chunkIndex < len(chunks); i++ {
		blockID := inode.I_block[i]
		if blockID == -1 {
			newBlock, err := fm.BlockManager.CreateBlock(chunks[chunkIndex])
			if err != nil {
				return fmt.Errorf("UpdateFile: Error creando bloque directo: %v", err)
			}
			inode.I_block[i] = newBlock
		} else {
			if err := fm.BlockManager.UpdateBlock(blockID, chunks[chunkIndex]); err != nil {
				return fmt.Errorf("UpdateFile: Error escribiendo bloque directo %d: %v", blockID, err)
			}
		}
		chunkIndex++
	}

	// Liberar bloques sobrantes directos
	for i := chunkIndex; i < 12; i++ {
		if inode.I_block[i] != -1 {
			_ = fm.BlockManager.DeleteBlock(inode.I_block[i])
			inode.I_block[i] = -1
		}
	}

	// Indirecto simple, doble y triple
	if chunkIndex < len(chunks) {
		if err := fm.writeIndirect(&inode, 12, 1, chunks[chunkIndex:], &chunkIndex); err != nil {
			return fmt.Errorf("UpdateFile: Error en indirecto simple: %v", err)
		}
	}
	if chunkIndex < len(chunks) {
		if err := fm.writeIndirect(&inode, 13, 2, chunks[chunkIndex:], &chunkIndex); err != nil {
			return fmt.Errorf("UpdateFile: Error en indirecto doble: %v", err)
		}
	}
	if chunkIndex < len(chunks) {
		if err := fm.writeIndirect(&inode, 14, 3, chunks[chunkIndex:], &chunkIndex); err != nil {
			return fmt.Errorf("UpdateFile: Error en indirecto triple: %v", err)
		}
	}

	// Actualizar tamaño y tiempo
	inode.I_size = 0
	for _, c := range chunks {
		inode.I_size += int64(len(c))
	}
	inode.I_mtime = time.Now().Unix()

	if err := fm.InodeManager.UpdateInode(inodeID, inode); err != nil {
		return fmt.Errorf("UpdateFile: Error guardando inodo actualizado: %v", err)
	}

	return nil
}
