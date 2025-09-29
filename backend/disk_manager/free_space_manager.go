package dmanager

import (
	disk "backend/disk_structs"
	"encoding/binary"
	"math"
	"sort"
)

type FreeSpaceManager struct {
}

type FreeSpace struct {
	Start int64
	End   int64
	Size  int64
}

func (fsm *FreeSpaceManager) GetFreeSpaces(mbr *disk.MBR) []FreeSpace {
	var spaces []FreeSpace
	var occupiedPartitions []disk.Partition

	for _, p := range mbr.Mbr_partition {
		if p.Part_status == '1' {
			occupiedPartitions = append(occupiedPartitions, p)
		}
	}

	sort.Slice(occupiedPartitions, func(i, j int) bool {
		return occupiedPartitions[i].Part_start < occupiedPartitions[j].Part_start
	})

	const MBR_SIZE = 256
	currentPos := int64(MBR_SIZE)

	for _, p := range occupiedPartitions {
		if p.Part_start > currentPos {
			size := p.Part_start - currentPos
			if size > 0 {
				spaces = append(spaces, FreeSpace{
					Start: currentPos,
					End:   p.Part_start - 1,
					Size:  size,
				})
			}
		}
		currentPos = p.Part_start + p.Part_size
	}

	if currentPos < mbr.Mbr_tamano {
		size := mbr.Mbr_tamano - currentPos
		if size > 0 {
			spaces = append(spaces, FreeSpace{
				Start: currentPos,
				End:   mbr.Mbr_tamano - 1,
				Size:  size,
			})
		}
	}

	return spaces
}

func (fsm *FreeSpaceManager) FindFirstFit(size int64, freeSpaces []FreeSpace) int64 {
	for _, space := range freeSpaces {
		if space.Size >= size {
			return space.Start
		}
	}
	return -1
}

func (fsm *FreeSpaceManager) FindBestFit(size int64, freeSpaces []FreeSpace) int64 {
	bestStart := int64(-1)
	minDifference := int64(math.MaxInt64)

	for _, space := range freeSpaces {
		if space.Size >= size {
			difference := space.Size - size
			if difference < minDifference {
				minDifference = difference
				bestStart = space.Start
			}
		}
	}
	return bestStart
}

func (fsm *FreeSpaceManager) FindWorstFit(size int64, freeSpaces []FreeSpace) int64 {
	var worstStart int64 = -1
	var maxSize int64 = -1

	for _, space := range freeSpaces {
		if space.Size >= size && space.Size > maxSize {
			maxSize = space.Size
			worstStart = space.Start
		}
	}
	return worstStart
}

func (fsm *FreeSpaceManager) GetFreeSpacesInExtended(extended disk.Partition, logicals []disk.EBR) []FreeSpace {
	var spaces []FreeSpace

	sort.Slice(logicals, func(i, j int) bool {
		return logicals[i].Part_start < logicals[j].Part_start
	})

	currentPos := extended.Part_start
	for _, l := range logicals {
		ebrStart := l.Part_start - int64(binary.Size(l))
		if ebrStart > currentPos {
			size := ebrStart - currentPos
			spaces = append(spaces, FreeSpace{Start: currentPos, End: ebrStart - 1, Size: size})
		}
		currentPos = l.Part_start + l.Part_size
	}

	if currentPos < (extended.Part_start + extended.Part_size) {
		size := (extended.Part_start + extended.Part_size) - currentPos
		spaces = append(spaces, FreeSpace{Start: currentPos, End: (extended.Part_start + extended.Part_size) - 1, Size: size})
	}

	return spaces
}
