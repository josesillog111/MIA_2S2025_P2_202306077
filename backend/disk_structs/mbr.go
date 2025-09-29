package disk

// atributos
type MBR struct {
	Mbr_tamano         int64        // size of the disk in bytes
	Mbr_fecha_creacion [20]byte     // Unix timestamp for creation date
	Mbr_disk_signature int64        // unique identifier for the disk
	Dsk_fit            byte         // 'B' = best fit, 'F' = first fit
	Mbr_partition      [4]Partition // array of partitions
}
