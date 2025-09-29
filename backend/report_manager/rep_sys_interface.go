package report

type ReportSysGenerator interface {
	Inode(path string, format string) error
	Block(path string, format string) error
	BitmapInode(path string, format string) error
	BitmapBlock(path string, format string) error
	Tree(path string, format string) error
	SuperBlock(path string, format string) error
	File(user_path string, virtualPath string, format string) error
	Ls(user_path string, virtualPath string, format string) error
}
