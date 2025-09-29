package report

type GenerateDiskReport interface {
	Mbr(path string, id string, format string) error
	Disk(path string, id string, format string) error
}
