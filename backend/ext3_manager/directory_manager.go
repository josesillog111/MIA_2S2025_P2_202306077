package ext3

import (
	dmanager "backend/disk_manager"
	file "backend/ext3_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type DirectoryManager struct {
	Partition    dmanager.MountedPartition
	SBManager    *SBManager
	InodeManager *InodeManager
	BlockManager *BlockManager
}

func NewDirectoryManager(partition dmanager.MountedPartition, sbm *SBManager) *DirectoryManager {
	return &DirectoryManager{
		Partition:    partition,
		SBManager:    sbm,
		InodeManager: NewInodeManager(partition, sbm),
		BlockManager: NewBlockManager(partition, sbm),
	}
}

func (dm *DirectoryManager) structToBytes(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (dm *DirectoryManager) bytesToStruct(b []byte, out interface{}) error {
	buf := bytes.NewReader(b)
	if err := binary.Read(buf, binary.LittleEndian, out); err != nil {
		return err
	}
	return nil
}

func (e *EXT3) removeRecursive(dirInodeID int64) error {
	entries, err := e.DirectoryManager.ReadDirectory(dirInodeID)
	if err != nil {
		return fmt.Errorf("RemoveRecursive: Error leyendo directorio %d: %v", dirInodeID, err)
	}

	for _, entry := range entries {
		name := strings.Trim(string(entry.B_name[:]), "\x00")
		childID := entry.B_inodo

		childInode, err := e.InodeManager.ReadInode(childID)
		if err != nil {
			return fmt.Errorf("RemoveRecursive: Error leyendo inodo hijo %d (%s): %v", childID, name, err)
		}

		// Verificar permiso de escritura antes de continuar
		if !e.hasPermission(childInode, "w") {
			return fmt.Errorf("RemoveRecursive: Permiso denegado en '%s'", name)
		}

		if childInode.I_type == FILE_TYPE {
			if err := e.FileManager.DeleteFile(dirInodeID, name); err != nil {
				return fmt.Errorf("RemoveRecursive: Error eliminando archivo '%s': %v", name, err)
			}
		} else if childInode.I_type == DIR_TYPE {
			if err := e.removeRecursive(childID); err != nil {
				return err // si falla uno, se cancela todo
			}
			if err := e.DirectoryManager.DeleteEntry(dirInodeID, name); err != nil {
				return fmt.Errorf("RemoveRecursive: Error eliminando subdirectorio '%s': %v", name, err)
			}
		}
	}

	// Finalmente eliminar el propio directorio (ya vacío)
	if err := e.DirectoryManager.DeleteDirectory(dirInodeID); err != nil {
		return fmt.Errorf("RemoveRecursive: Error eliminando directorio %d: %v", dirInodeID, err)
	}

	return nil
}

/*
	CRUD MANEJO DE LAS ESTRUCURAS
	"DIRECTORIO". TAMBIÉN SE MODIFICAN LOS BITMAPS
*/

func (dm *DirectoryManager) CreateDirectory(parentInodeID int64, name string, owner string, perms uint16) (int64, error) {
	// 1. Crear un inodo de tipo directorio
	inodeID, inode, err := dm.InodeManager.CreateInode(DIR_TYPE, 0, owner, perms)
	if err != nil {
		return -1, fmt.Errorf("CreateDirectory: error creando inodo: %v", err)
	}

	// 2. Inicializar bloque de directorio (con entradas vacías)
	var dir file.Directory
	for i := range dir.B_content {
		dir.B_content[i].B_inodo = -1
	}

	// 3. Escribir bloque en el primer puntero del inodo
	data, err := dm.structToBytes(&dir)
	if err != nil {
		_ = dm.InodeManager.DeleteInode(inodeID) // rollback
		return -1, fmt.Errorf("CreateDirectory: error serializando bloque: %v", err)
	}

	// Reservar el bloque dentro del inodo
	blockID := inode.I_block[0] // suponemos que el primer bloque es 0 inicialmente
	if err := dm.BlockManager.UpdateBlock(blockID, data); err != nil {
		_ = dm.InodeManager.DeleteInode(inodeID) // rollback
		return -1, fmt.Errorf("CreateDirectory: error escribiendo bloque: %v", err)
	}

	// 4. Registrar la entrada en el directorio padre
	if err := dm.AddEntry(parentInodeID, name, inodeID); err != nil {
		_ = dm.InodeManager.DeleteInode(inodeID) // rollback
		return -1, fmt.Errorf("CreateDirectory: error agregando entrada al directorio padre: %v", err)
	}

	return inodeID, nil
}

func (dm *DirectoryManager) ReadDirectory(dirInodeID int64) ([]file.DirEntry, error) {

	inode, err := dm.InodeManager.ReadInode(dirInodeID)
	if err != nil {
		return nil, fmt.Errorf("ReadDirectory: error leyendo inodo %d: %v", dirInodeID, err)
	}

	if inode.I_type != DIR_TYPE {
		return nil, fmt.Errorf("ReadDirectory: inodo %d no es un directorio", dirInodeID)
	}

	var entries []file.DirEntry

	// Recorremos los bloques directos del inodo
	for _, blockID := range inode.I_block {
		if blockID == -1 {
			continue
		}

		raw, err := dm.BlockManager.ReadBlock(blockID)
		if err != nil {
			return nil, fmt.Errorf("ReadDirectory: error leyendo bloque %d: %v", blockID, err)
		}

		var dir file.Directory
		buf := bytes.NewReader(raw)
		if err := binary.Read(buf, binary.LittleEndian, &dir); err != nil {
			return nil, fmt.Errorf("ReadDirectory: error parseando bloque %d: %v", blockID, err)
		}

		for _, entry := range dir.B_content {
			name := strings.Trim(string(entry.B_name[:]), "\x00")
			if name != "" && entry.B_inodo != -1 {
				entries = append(entries, entry)
			}
		}
	}

	return entries, nil
}

func (dm *DirectoryManager) UpdateDirectory(dirInodeID int64, dir file.Directory) error {
	// 1. Leer inodo del directorio
	inode, err := dm.InodeManager.ReadInode(dirInodeID)
	if err != nil {
		return fmt.Errorf("UpdateDirectory: error leyendo inodo %d: %v", dirInodeID, err)
	}

	if byte(inode.I_type) != DIR_TYPE {
		return fmt.Errorf("UpdateDirectory: inodo %d no es un directorio", dirInodeID)
	}

	// 2. Serializar la estructura de directorio
	data, err := dm.structToBytes(&dir)
	if err != nil {
		return fmt.Errorf("UpdateDirectory: error serializando directorio: %v", err)
	}

	// 3. Escribir en el primer bloque del inodo (puedes expandir a múltiples bloques si es necesario)
	blockID := inode.I_block[0]
	if blockID == -1 {
		// No hay bloque asignado aún, asignamos uno
		newBlockID, err := dm.BlockManager.AllocateBlock()
		if err != nil {
			return fmt.Errorf("UpdateDirectory: sin bloques libres: %v", err)
		}
		blockID = newBlockID
		inode.I_block[0] = blockID

		// Actualizar inodo con el nuevo bloque asignado
		if err := dm.InodeManager.UpdateInode(dirInodeID, inode); err != nil {
			_ = dm.BlockManager.DeallocateBlock(blockID)
			return fmt.Errorf("UpdateDirectory: error actualizando inodo: %v", err)
		}
	}

	// 4. Escribir datos en bloque
	if err := dm.BlockManager.UpdateBlock(blockID, data); err != nil {
		return fmt.Errorf("UpdateDirectory: error escribiendo bloque %d: %v", blockID, err)
	}

	return nil
}

func (dm *DirectoryManager) DeleteDirectory(dirInodeID int64) error {
	// 1. Leer entradas
	entries, err := dm.ReadDirectory(dirInodeID)
	if err != nil {
		return err
	}

	// 2. Asegurarse que está vacío
	if len(entries) > 0 {
		return fmt.Errorf("DeleteDirectory: directorio no vacío")
	}

	// 3. Liberar inodo (DeleteInode ya se encarga de liberar el bitmap y limpiar bloques)
	if err := dm.InodeManager.DeleteInode(dirInodeID); err != nil {
		return fmt.Errorf("DeleteDirectory: error eliminando inodo: %v", err)
	}

	return nil
}

/*
func (dm *DirectoryManager) DeleteDirectoryRecursive(dirInodeID int64) error {
	entries, err := dm.ReadDirectory(dirInodeID)
	if err != nil {
		return err
	}

	for _, e := range entries {
		childInode := e.B_inodo
		name := strings.Trim(string(e.B_name[:]), "\x00")
		childInodeStruct, err := dm.InodeManager.ReadInode(childInode)
		if err != nil {
			return err
		}
		if childInodeStruct.I_type == DIR_TYPE {
			if err := dm.DeleteDirectoryRecursive(childInode); err != nil {
				return err
			}
		} else {
			if err := dm.InodeManager.DeleteInode(childInode); err != nil {
				return err
			}
		}
	}
	return dm.DeleteDirectory(dirInodeID)
}
*/

func (dm *DirectoryManager) AddEntry(parentInodeID int64, name string, inodeID int64) error {
	// 1. Validar ID padre
	if parentInodeID < 0 {
		return fmt.Errorf("AddEntry: parentInodeID inválido (%d)", parentInodeID)
	}

	// 2. Leer inodo padre
	inode, err := dm.InodeManager.ReadInode(parentInodeID)
	if err != nil {
		return fmt.Errorf("AddEntry: error leyendo inodo padre %d: %v", parentInodeID, err)
	}

	// 3. Verificar que sea directorio
	type_ino := byte(inode.I_type)
	if type_ino != DIR_TYPE {
		return fmt.Errorf("AddEntry: inodo %d no es un directorio (type_ino=%v, DIR_TYPE=%v)", parentInodeID, type_ino, DIR_TYPE)
	}

	// 4. Buscar bloque con espacio
	for _, bid := range inode.I_block {
		if bid == -1 {
			continue
		}
		data, err := dm.BlockManager.ReadBlock(bid)
		if err != nil {
			return fmt.Errorf("AddEntry: error leyendo bloque %d: %v", bid, err)
		}

		var dir file.Directory
		if err := dm.bytesToStruct(data, &dir); err != nil {
			return fmt.Errorf("AddEntry: error parseando bloque %d: %v", bid, err)
		}

		// 5. Buscar slot vacío
		for i := range dir.B_content {
			if dir.B_content[i].B_inodo == -1 {
				copy(dir.B_content[i].B_name[:], []byte(name))
				dir.B_content[i].B_inodo = inodeID

				db, err := dm.structToBytes(&dir)
				if err != nil {
					return fmt.Errorf("AddEntry: error serializando bloque actualizado: %v", err)
				}
				if err := dm.BlockManager.UpdateBlock(bid, db); err != nil {
					return fmt.Errorf("AddEntry: error actualizando bloque %d: %v", bid, err)
				}

				return nil
			}
		}
	}

	// 6. Si no hay espacio → reservar nuevo bloque
	newBlock, err := dm.BlockManager.AllocateBlock()
	if err != nil {
		return fmt.Errorf("AddEntry: no hay bloques libres")
	}

	// Inicializar bloque de directorio
	var dir file.Directory
	for i := range dir.B_content {
		dir.B_content[i].B_inodo = -1
	}
	copy(dir.B_content[0].B_name[:], []byte(name))
	dir.B_content[0].B_inodo = inodeID

	db, err := dm.structToBytes(&dir)
	if err != nil {
		_ = dm.BlockManager.DeallocateBlock(newBlock)
		return fmt.Errorf("AddEntry: error serializando nuevo bloque: %v", err)
	}
	if err := dm.BlockManager.UpdateBlock(newBlock, db); err != nil {
		_ = dm.BlockManager.DeallocateBlock(newBlock)
		return fmt.Errorf("AddEntry: error escribiendo nuevo bloque: %v", err)
	}

	// 7. Asignar nuevo bloque al inodo padre
	assigned := false
	for i := range inode.I_block {
		if inode.I_block[i] == -1 {
			inode.I_block[i] = newBlock
			assigned = true
			break
		}
	}
	if !assigned {
		_ = dm.BlockManager.DeallocateBlock(newBlock)
		return fmt.Errorf("AddEntry: no hay punteros libres en el inodo del padre")
	}

	// 8. Actualizar tamaño y guardar inodo padre
	inode.I_size += dm.SBManager.SB.S_block_size
	if err := dm.InodeManager.UpdateInode(parentInodeID, inode); err != nil {
		return fmt.Errorf("AddEntry: error actualizando inodo padre: %v", err)
	}

	return nil
}

func (dm *DirectoryManager) DeleteEntry(parentInodeID int64, name string) error {
	// 1️⃣ Leer el inodo del directorio padre
	inode, err := dm.InodeManager.ReadInode(parentInodeID)
	if err != nil {
		return fmt.Errorf("DeleteEntry: error leyendo inodo padre %d: %v", parentInodeID, err)
	}
	if inode.I_type != DIR_TYPE {
		return fmt.Errorf("DeleteEntry: el inodo %d no es un directorio", parentInodeID)
	}

	// 2️⃣ Buscar en los bloques directos del directorio
	for _, blockID := range inode.I_block {
		if blockID == -1 {
			continue
		}

		raw, err := dm.BlockManager.ReadBlock(blockID)
		if err != nil {
			return fmt.Errorf("DeleteEntry: error leyendo bloque %d: %v", blockID, err)
		}

		var dir file.Directory
		buf := bytes.NewReader(raw)
		if err := binary.Read(buf, binary.LittleEndian, &dir); err != nil {
			return fmt.Errorf("DeleteEntry: error parseando bloque %d: %v", blockID, err)
		}

		// 3️⃣ Buscar el nombre dentro de las 4 entradas
		updated := false
		for i := range dir.B_content {
			entryName := strings.Trim(string(dir.B_content[i].B_name[:]), "\x00")
			if entryName == name {
				dir.B_content[i].B_inodo = -1
				for j := range dir.B_content[i].B_name {
					dir.B_content[i].B_name[j] = 0
				}
				updated = true
				break
			}
		}

		// 4️⃣ Si se encontró y modificó → reescribir el bloque
		if updated {
			bufOut := new(bytes.Buffer)
			if err := binary.Write(bufOut, binary.LittleEndian, &dir); err != nil {
				return fmt.Errorf("DeleteEntry: error serializando bloque %d: %v", blockID, err)
			}
			if err := dm.BlockManager.UpdateBlock(blockID, bufOut.Bytes()); err != nil {
				return fmt.Errorf("DeleteEntry: error actualizando bloque %d: %v", blockID, err)
			}
			return nil
		}
	}

	// 5️⃣ Si no se encontró la entrada
	return fmt.Errorf("DeleteEntry: '%s' no encontrado en el directorio padre", name)
}

func (dm *DirectoryManager) RenameEntry(parentInodeID int64, oldName string, newName string) error {
	inode, err := dm.InodeManager.ReadInode(parentInodeID)
	if err != nil {
		return fmt.Errorf("RenameEntry: Error leyendo inodo padre: %v", err)
	}
	if inode.I_type != DIR_TYPE {
		return fmt.Errorf("RenameEntry: El inodo %d no es un directorio", parentInodeID)
	}

	for _, blockID := range inode.I_block {
		if blockID == -1 {
			continue
		}

		raw, err := dm.BlockManager.ReadBlock(blockID)
		if err != nil {
			return fmt.Errorf("RenameEntry: Error leyendo bloque %d: %v", blockID, err)
		}

		var dir file.Directory
		buf := bytes.NewReader(raw)
		if err := binary.Read(buf, binary.LittleEndian, &dir); err != nil {
			return fmt.Errorf("RenameEntry: Error parseando bloque %d: %v", blockID, err)
		}

		updated := false
		for i := range dir.B_content {
			entryName := strings.Trim(string(dir.B_content[i].B_name[:]), "\x00")
			if entryName == oldName {
				// Sobrescribir nombre
				for j := range dir.B_content[i].B_name {
					dir.B_content[i].B_name[j] = 0
				}
				copy(dir.B_content[i].B_name[:], []byte(newName))
				updated = true
				break
			}
		}

		if updated {
			bufOut := new(bytes.Buffer)
			if err := binary.Write(bufOut, binary.LittleEndian, &dir); err != nil {
				return fmt.Errorf("RenameEntry: Error serializando bloque: %v", err)
			}
			if err := dm.BlockManager.UpdateBlock(blockID, bufOut.Bytes()); err != nil {
				return fmt.Errorf("RenameEntry: Error actualizando bloque: %v", err)
			}
			return nil
		}
	}

	return fmt.Errorf("RenameEntry: No se encontró '%s' en el directorio padre", oldName)
}

func (dm *DirectoryManager) ResolvePath(path string) (int64, string, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	// Tomar el inodo raíz desde el superbloque (esto evita desfases manuales)
	sb, err := dm.SBManager.ReadSuperBlock()
	if err != nil {
		return -1, "", fmt.Errorf("ResolvePath: error leyendo superbloque: %v", err)
	}
	currentID := sb.S_first_ino // raíz siempre es el primer inodo
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		nextID, err := dm.FindEntry(currentID, part)
		if err != nil {
			return -1, "", fmt.Errorf("directorio no existe: %s", part)
		}
		currentID = nextID
	}

	last := parts[len(parts)-1]
	return currentID, last, nil
}

func (dm *DirectoryManager) FindEntry(dirInodeID int64, name string) (int64, error) {
	entries, err := dm.ReadDirectory(dirInodeID)
	if err != nil {
		return -1, fmt.Errorf("FindEntry: error leyendo directorio %d: %v", dirInodeID, err)
	}

	for _, entry := range entries {
		entryName := strings.Trim(string(entry.B_name[:]), "\x00")
		if entryName == name {
			return entry.B_inodo, nil
		}
	}
	return -1, fmt.Errorf("FindEntry: no existe entrada con nombre '%s' en dir %d", name, dirInodeID)
}

func (dm *DirectoryManager) ListDirectory(dirInodeID int64) ([]string, error) {
	entries, err := dm.ReadDirectory(dirInodeID)
	if err != nil {
		return nil, fmt.Errorf("ListDirectory: error leyendo directorio %d: %v", dirInodeID, err)
	}

	var names []string
	for _, entry := range entries {
		names = append(names, strings.Trim(string(entry.B_name[:]), "\x00"))
	}
	return names, nil
}
