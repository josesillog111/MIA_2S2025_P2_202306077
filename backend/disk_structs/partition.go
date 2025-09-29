package disk

import (
	"fmt"
	"strings"
)

type Partition struct {
	Part_status      byte     // 1 = mounted, 0 = unmounted
	Part_type        byte     // 'P' = primary, 'E' = extended
	Part_fit         byte     // 'B' = best fit, 'F' = first fit, W = worst fit
	_                byte     // padding
	Part_start       int64    // start position of the partition
	Part_size        int64    // size of the partition in bytes
	Part_name        [16]byte // name of the partition
	Part_correlative int64    // correlation ID for the partition (-1 if not mounted, else start in 1 )
	Part_id          [4]byte  // unique identifier for the partition
}

func (p Partition) ToString() string {
	return fmt.Sprintf("{Status: %c, Type: %c, Fit: %c, Start: %d, Size: %d, Name: %s, Correlative: %d, Id: %s}",
		p.Part_status,
		p.Part_type,
		p.Part_fit,
		p.Part_start,
		p.Part_size,
		strings.Trim(string(p.Part_name[:]), "\x00"),
		p.Part_correlative,
		strings.Trim(string(p.Part_id[:]), "\x00"),
	)
}
