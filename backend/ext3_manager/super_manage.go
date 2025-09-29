package ext3

import (
	dmanager "backend/disk_manager"
	file "backend/ext3_structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type SBManager struct {
	SB        file.SuperBlock
	Partition dmanager.MountedPartition
}

func NewSBManager(partition dmanager.MountedPartition) *SBManager {
	return &SBManager{
		Partition: partition,
	}
}

func (sbm *SBManager) UpdateLastMountedDate() error {
	var mtimeBytes [20]byte
	binary.LittleEndian.PutUint64(mtimeBytes[:8], uint64(time.Now().Unix()))
	sbm.SB.S_mtime = mtimeBytes
	err := sbm.WriteSuperBlock(sbm.SB)
	return err
}

func (sbm *SBManager) UpdateLastUnmountedDate() error {
	var mtimeBytes [20]byte
	binary.LittleEndian.PutUint64(mtimeBytes[:8], uint64(time.Now().Unix()))
	sbm.SB.S_umtime = mtimeBytes
	err := sbm.WriteSuperBlock(sbm.SB)
	return err
}

func (sbm *SBManager) UpdateMountedCount() error {
	sbm.SB.S_mnt_count++
	err := sbm.WriteSuperBlock(sbm.SB)
	return err
}

func (sbm *SBManager) ReadSuperBlock() (file.SuperBlock, error) {
	var sb file.SuperBlock
	disk, err := os.OpenFile(sbm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return sb, fmt.Errorf("ReadSuperBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()
	offset := sbm.Partition.Partition.Part_start
	size := int64(512)
	if int64(size) > sbm.Partition.Partition.Part_size {
		return sb, fmt.Errorf("ReadSuperBlock: El superbloque no cabe en la partición")
	}

	buf := make([]byte, size)
	if _, err := disk.ReadAt(buf, offset); err != nil {
		return sb, fmt.Errorf("ReadSuperBlock: Error leyendo superblock: %v", err)
	}

	if err := binary.Read(bytes.NewReader(buf), binary.LittleEndian, &sb); err != nil {
		return sb, fmt.Errorf("ReadSuperBlock: Error deserializando superblock: %v", err)
	}

	sbm.SB = sb

	return sb, nil
}

func (sbm *SBManager) WriteSuperBlock(block file.SuperBlock) error {
	disk, err := os.OpenFile(sbm.Partition.Path, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("WriteSuperBlock: Error abriendo disco: %v", err)
	}
	defer disk.Close()

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, block); err != nil {
		return fmt.Errorf("WriteSuperBlock: Error serializando superblock: %v", err)
	}

	offset := sbm.Partition.Partition.Part_start
	if int64(len(buf.Bytes())) > sbm.Partition.Partition.Part_size {
		return fmt.Errorf("WriteSuperBlock: El superbloque no cabe en la partición")
	}

	if _, err := disk.WriteAt(buf.Bytes(), offset); err != nil {
		return fmt.Errorf("WriteSuperBlock: Error escribiendo superblock: %v", err)
	}
	sbm.SB = block
	return nil
}

func (sbm *SBManager) NewSuperBlock(offset_partition int64, partition_size int64, file_sys [10]byte) (*file.SuperBlock, int64, int64) {
	var sb file.SuperBlock
	sb, bm_inode_size, bm_block_size := file.NewSuperBlock(offset_partition, partition_size, file_sys)
	return &sb, bm_inode_size, bm_block_size
}
