package file_ext3

type Information struct {
	I_operation [10]byte
	I_path      [32]byte
	I_content   [64]byte
	I_date      [20]byte
}
