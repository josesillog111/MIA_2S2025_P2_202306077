package file

import "fmt"

type Inode struct {
	I_uid         int64
	I_gid         int64
	I_size        int64
	I_atime       int64     // time of last access
	I_ctime       int64     // time of creation
	I_mtime       int64     // time of last modification
	I_block       [15]int64 //array de 15 bloques
	I_type        byte      // 1 = file, 0 = dir
	I_permissions uint16
}

func NewInode(fileType byte, size int64, perms int64, createDate string) Inode {
	var I_ctime [20]byte
	copy(I_ctime[:], createDate)

	return Inode{
		I_uid:         1,
		I_gid:         1,
		I_size:        size,
		I_atime:       0,
		I_ctime:       0,
		I_mtime:       0,
		I_block:       [15]int64{},
		I_type:        fileType,
		I_permissions: uint16(perms),
	}
}

func (ino *Inode) ToString() string {
	return fmt.Sprintf("Inode{UID: %d, GID: %d, Size: %d, Type: %d, Permissions: %o, Owner: %d}",
		ino.I_uid, ino.I_gid, ino.I_size, ino.I_type, ino.I_permissions, ino.I_uid)
}
