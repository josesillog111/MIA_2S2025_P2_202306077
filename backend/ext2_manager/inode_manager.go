package ext2

import (
	dmanager "backend/disk_manager"
	file "backend/ext2_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
)

type InodeManager struct {
	Partition dmanager.MountedPartition
	SBManager *SBManager
}

func NewInodeManager(partition dmanager.MountedPartition, sbm *SBManager) *InodeManager {
	return &InodeManager{
		Partition: partition,
		SBManager: sbm,
	}
}

/* 		CRUD Inodes 	*/

// FileType:: 0 = file, 1 = directory

// Crear inodo
func (im *InodeManager) CreateInode(fileType byte, size int64, owner string, perms uint16) (int64, file.Inode, error) {
	// 1. Obtener un inodo libre del bitmap
	inodeID, err := im.AllocateInode()
	if err != nil {
		return -1, file.Inode{}, fmt.Errorf("CreateInode: %v", err)
	}

	// 2. Construir el inodo
	var inode file.Inode
	inode.I_uid = 1
	inode.I_gid = 1
	inode.I_type = fileType
	inode.I_size = size
	inode.I_ctime = time.Now().Unix()
	inode.I_atime = inode.I_ctime
	inode.I_mtime = inode.I_ctime

	for i := range inode.I_block {
		inode.I_block[i] = -1
	}
	inode.I_permissions = perms

	// 3. Guardar inodo en disco
	if err := im.UpdateInode(inodeID, inode); err != nil {
		// Si falla, liberar inodo
		_ = im.DeleteInode(inodeID)
		return -1, file.Inode{}, fmt.Errorf("CreateInode: Error escribiendo inodo: %v", err)
	}
	return inodeID, inode, nil
}

// Leer inodo
func (im *InodeManager) ReadInode(id int64) (file.Inode, error) {
	sb, err := im.SBManager.ReadSuperBlock()
	if err != nil {
		return file.Inode{}, fmt.Errorf("ReadInode: Error leyendo superblock: %v", err)
	}
	if id < 0 || id >= sb.S_inodes_count {
		return file.Inode{}, fmt.Errorf("ReadInode: Id fuera de rango (%d)", id)
	}

	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return file.Inode{}, fmt.Errorf("ReadInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	offset := sb.S_inode_start + id*sb.S_inode_size

	buf := make([]byte, sb.S_inode_size)
	if _, err := disk.ReadAt(buf, offset); err != nil {
		return file.Inode{}, fmt.Errorf("ReadInode: Error leyendo inodo %d: %v", id, err)
	}

	bufReader := bytes.NewReader(buf)
	var inode file.Inode

	// Deserialización
	if err := binary.Read(bufReader, binary.LittleEndian, &inode.I_uid); err != nil {
		return inode, fmt.Errorf("ReadInode: Read I_uid: %v", err)
	}
	binary.Read(bufReader, binary.LittleEndian, &inode.I_gid)
	binary.Read(bufReader, binary.LittleEndian, &inode.I_size)
	binary.Read(bufReader, binary.LittleEndian, &inode.I_atime)
	binary.Read(bufReader, binary.LittleEndian, &inode.I_ctime)
	binary.Read(bufReader, binary.LittleEndian, &inode.I_mtime)

	for i := range inode.I_block {
		binary.Read(bufReader, binary.LittleEndian, &inode.I_block[i])
	}

	binary.Read(bufReader, binary.LittleEndian, &inode.I_type)
	binary.Read(bufReader, binary.LittleEndian, &inode.I_permissions)

	return inode, nil
}

// Eliminar inodo
func (im *InodeManager) DeleteInode(id int64) error {
	sb := im.SBManager.SB
	if id < 0 || id >= sb.S_inodes_count {
		return fmt.Errorf("DeleteInode: Id fuera de rango (%d)", id)
	}

	// Liberar inodo en bitmap y actualizar superblock
	if err := im.DeallocateInode(id); err != nil {
		return fmt.Errorf("DeleteInode: Error liberando inodo: %v", err)
	}

	// Sobrescribir inodo con ceros
	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeleteInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	empty := make([]byte, sb.S_inode_size)
	offset := sb.S_inode_start + id*sb.S_inode_size
	if _, err := disk.WriteAt(empty, offset); err != nil {
		return fmt.Errorf("DeleteInode: Error limpiando inodo: %v", err)
	}

	return nil
}

// Actualizar inodo
func (im *InodeManager) UpdateInode(id int64, inode file.Inode) error {
	sb, err := im.SBManager.ReadSuperBlock()
	if err != nil {
		return fmt.Errorf("UpdateInode: Error leyendo superblock: %v", err)
	}

	if id < 0 || id >= sb.S_inodes_count {
		return fmt.Errorf("UpdateInode: Id fuera de rango (%d)", id)
	}

	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("UpdateInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	buf := new(bytes.Buffer)

	// Serialización manual consistente y comprobando errores
	if err := binary.Write(buf, binary.LittleEndian, inode.I_uid); err != nil {
		return fmt.Errorf("UpdateInode: Write I_uid: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_gid); err != nil {
		return fmt.Errorf("UpdateInode: Write I_gid: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_size); err != nil {
		return fmt.Errorf("UpdateInode: Write I_size: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_atime); err != nil {
		return fmt.Errorf("UpdateInode: Write I_atime: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_ctime); err != nil {
		return fmt.Errorf("UpdateInode: Write I_ctime: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_mtime); err != nil {
		return fmt.Errorf("UpdateInode: Write I_mtime: %v", err)
	}
	for _, b := range inode.I_block {
		if err := binary.Write(buf, binary.LittleEndian, b); err != nil {
			return fmt.Errorf("UpdateInode: Write I_block entry: %v", err)
		}
	}

	// Escribir I_type e I_permissions con binary.Write para mantener consistencia
	if err := binary.Write(buf, binary.LittleEndian, inode.I_type); err != nil {
		return fmt.Errorf("UpdateInode: Write I_type: %v", err)
	}
	if err := binary.Write(buf, binary.LittleEndian, inode.I_permissions); err != nil {
		return fmt.Errorf("UpdateInode: Write I_permissions: %v", err)
	}

	// Rellenar padding hasta completar el tamaño del inodo
	expected := int(sb.S_inode_size)
	padding := expected - buf.Len()
	if padding < 0 {
		return fmt.Errorf("UpdateInode: Tamaño serializado (%d) excede S_inode_size (%d)", buf.Len(), expected)
	}
	if padding > 0 {
		buf.Write(make([]byte, padding))
	}

	offset := sb.S_inode_start + id*sb.S_inode_size
	if _, err := disk.WriteAt(buf.Bytes(), offset); err != nil {
		return fmt.Errorf("UpdateInode: Error escribiendo inodo %d: %v", id, err)
	}

	return nil
}

/*  OPERACIONES DE BAJO NIVEL SOBRE EL BITMAP  */

// Asignar inodo
func (im *InodeManager) AllocateInode() (int64, error) {
	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return -1, fmt.Errorf("AllocateInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	sb := im.SBManager.SB

	// Calcular tamaño del bitmap en bytes
	bitmapSize := (sb.S_inodes_count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	// Leer bitmap de inodos
	bitmapStart := sb.S_bm_inode_start
	if _, err := disk.ReadAt(bitmap, bitmapStart); err != nil {
		return -1, fmt.Errorf("AllocateInode: Error leyendo bitmap de inodos: %v", err)
	}

	// Buscar un inodo libre (0 = libre, 1 = ocupado)
	for byteIndex := 0; byteIndex < len(bitmap); byteIndex++ {
		for bit := 0; bit < 8; bit++ {
			inodeID := int64(byteIndex*8 + bit)

			// Saltar IDs menores al primer inodo válido
			if inodeID < sb.S_first_ino {
				continue
			}
			if inodeID >= sb.S_inodes_count {
				break
			}

			if bitmap[byteIndex]&(1<<bit) == 0 {
				// Marcar como ocupado
				bitmap[byteIndex] |= 1 << bit

				// Escribir bitmap actualizado
				if _, err := disk.WriteAt(bitmap, bitmapStart); err != nil {
					return -1, fmt.Errorf("AllocateInode: Error actualizando bitmap de inodos: %v", err)
				}

				// Actualizar superbloque
				sb.S_free_inodes_count--
				im.SBManager.SB = sb
				if err := im.SBManager.WriteSuperBlock(sb); err != nil {
					return -1, fmt.Errorf("AllocateInode: Error escribiendo superbloque: %v", err)
				}

				return inodeID, nil
			}
		}
	}

	return -1, fmt.Errorf("AllocateInode: No hay inodos libres")
}

// Liberar inodo
func (im *InodeManager) DeallocateInode(id int64) error {
	sb := im.SBManager.SB
	if id < sb.S_first_ino || id >= sb.S_inodes_count {
		return fmt.Errorf("DeallocateInode: Id de inodo fuera de rango o reservado: %d", id)
	}

	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("DeallocateInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	// Calcular tamaño del bitmap en bytes
	bitmapSize := (sb.S_inodes_count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	// Leer bitmap de inodos
	if _, err := disk.ReadAt(bitmap, sb.S_bm_inode_start); err != nil {
		return fmt.Errorf("DeallocateInode: Error leyendo bitmap de inodos: %v", err)
	}

	// Liberar bit correspondiente
	byteIndex := id / 8
	bitIndex := id % 8
	bitmap[byteIndex] &^= 1 << bitIndex

	// Escribir bitmap actualizado
	if _, err := disk.WriteAt(bitmap, sb.S_bm_inode_start); err != nil {
		return fmt.Errorf("DeallocateInode: Error escribiendo bitmap de inodos: %v", err)
	}

	// Actualizar superbloque
	sb.S_free_inodes_count++
	im.SBManager.SB = sb
	if err := im.SBManager.WriteSuperBlock(sb); err != nil {
		return fmt.Errorf("DeallocateInode: Error actualizando superbloque: %v", err)
	}

	return nil
}

/*

INICIALIZAR BITMAP

*/

// Inicializar bitmap
func (im *InodeManager) InitInodeBitmap(offset int64, count int64, firstIno int64) error {
	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("InitInodeBitmap: No se pudo abrir el disco %s: %v", im.Partition.Path, err)
	}
	defer disk.Close()

	realOffset := im.Partition.Partition.Part_start + offset

	bitmapSize := (count + 7) / 8
	buf := make([]byte, int(bitmapSize)) // todo en 0 = libre

	if _, err := disk.WriteAt(buf, realOffset); err != nil {
		return fmt.Errorf("InitInodeBitmap: Error al escribir bitmap de inodos: %v", err)
	}

	return nil
}

func (im *InodeManager) GetBitmapInode(offset int64, count int64) (string, error) {
	disk, err := os.OpenFile(im.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return "", fmt.Errorf("GetBitmapInode: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	realOffset := im.SBManager.SB.S_bm_inode_start
	bitmapSize := (count + 7) / 8
	bitmap := make([]byte, bitmapSize)

	if _, err := disk.ReadAt(bitmap, realOffset); err != nil {
		return "", fmt.Errorf("GetBitmapInode: Error leyendo bitmap de inodos: %v", err)
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
