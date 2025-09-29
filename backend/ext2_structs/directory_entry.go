package file

type DirEntry struct {
	B_name  [12]byte // Name of the content
	B_inodo int64    // Inode number associated with the content
}
