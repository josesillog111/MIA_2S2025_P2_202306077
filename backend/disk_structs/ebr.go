package disk

type EBR struct {
	Part_mount byte     // 1 = mounted, 0 = unmounted
	Part_fit   byte     // 'B' = best fit, 'F' = first fit
	Part_start int64    // start position of the extended partition
	Part_size  int64    // size of the extended partition in bytes
	Part_next  int64    // start position of the next EBR in the chain
	Part_name  [16]byte // name of the extended partition
}
