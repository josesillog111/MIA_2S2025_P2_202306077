package file_ext3

import "fmt"

type SuperBlock struct {
	S_filesystem_type   [10]byte // Type of filesystem (e.g., 0 for ext4, 1 for NTFS)
	S_inodes_count      int64    // Total number of inodes
	S_blocks_count      int64    // Total number of blocks
	S_free_blocks_count int64    // Number of free blocks
	S_free_inodes_count int64    // Number of free inodes
	S_mtime             [20]byte // Last mount time
	S_umtime            [20]byte // Last unmount time
	S_mnt_count         int64    // Mount count
	S_magic             int64    // Magic number to identify the filesystem type
	S_inode_size        int64    // Size of an inode
	S_block_size        int64    // Size of a block
	S_first_ino         int64    // First inode number
	S_first_blo         int64    // First block number
	S_bm_inode_start    int64    // Start position of the bitmap for inodes
	S_bm_block_start    int64    // Start position of the bitmap for blocks
	S_inode_start       int64    // Start position of the inode table
	S_block_start       int64    // Start position of the block table
}

func NewSuperBlock(partition_start, partition_size int64, file_sys [10]byte) (SuperBlock, int64, int64) {
	sb_size := int64(512)
	block_size := int64(512)
	inode_size := int64(256) // más realista, en EXT2 son 128 bytes

	// número máximo de inodos = partición / 4KB (heurística)
	totalInodes := partition_size / (4 * 1024)
	if totalInodes < 16 {
		totalInodes = 16
	}
	totalBlocks := (partition_size - sb_size) / block_size

	bm_inode_size := (totalInodes + 7) / 8
	bm_block_size := (totalBlocks + 7) / 8

	sb := SuperBlock{
		S_filesystem_type:   file_sys,
		S_inodes_count:      totalInodes,
		S_blocks_count:      totalBlocks,
		S_free_blocks_count: totalBlocks,
		S_free_inodes_count: totalInodes,
		S_mtime:             [20]byte{},
		S_umtime:            [20]byte{},
		S_mnt_count:         0,
		S_magic:             0xEF53, // valor típico de EXT2/EXT3/EXT4
		S_inode_size:        inode_size,
		S_block_size:        block_size,
		S_first_ino:         1,
		S_first_blo:         1,
		S_bm_inode_start:    partition_start + sb_size,
		S_bm_block_start:    partition_start + sb_size + bm_inode_size,
		S_inode_start:       partition_start + sb_size + bm_inode_size + bm_block_size,
		S_block_start:       partition_start + sb_size + bm_inode_size + bm_block_size + totalInodes*inode_size,
	}

	return sb, bm_inode_size, bm_block_size
}

func (sb *SuperBlock) ToString() string {
	return fmt.Sprintf(`{
		Filesystem Type: %s, 
		Inodes: %d,
		Blocks: %d,
		Free Blocks: %d,
		Free Inodes: %d,
		Block Size: %d,
		Inode Size: %d,
		BM Inode Start: %d,
		BM Block Start: %d,
		Inode Start: %d,
		Block Start: %d,
		Magic: %d,
		Mount Count: %d,
		Last Mount Time: %s,
		Last Unmount Time: %s,
		First Inode: %d,
		First Block: %d
	}`, string(sb.S_filesystem_type[:]), sb.S_inodes_count, sb.S_blocks_count, sb.S_free_blocks_count, sb.S_free_inodes_count, sb.S_block_size, sb.S_inode_size, sb.S_bm_inode_start, sb.S_bm_block_start, sb.S_inode_start, sb.S_block_start, sb.S_magic, sb.S_mnt_count, string(sb.S_mtime[:]), string(sb.S_umtime[:]), sb.S_first_ino, sb.S_first_blo)
}
