package gid

import "errors"

const TotalBits = 1 << 6

type BitsAllocator struct {
	timestampBits uint32
	workerIdBits  uint32
	sequenceBits  uint32

	/**
	 * Max value for workId & sequence
	 */
	maxDeltaSeconds int64
	maxWorkerId     int64
	maxSequence     int64

	/**
	 * Shift for timestamp & workerId
	 */
	timestampShift uint32
	workerIdShift  uint32
}

func NewBitsAllocator(timestampBits, workerIdBits, sequenceBits uint32) *BitsAllocator {
	var signBits uint32 = 1

	allocateTotalBits := signBits + timestampBits + workerIdBits + sequenceBits

	if allocateTotalBits > TotalBits {
		panic(errors.New("allocate larger than 64 bits"))
	}

	return &BitsAllocator{
		timestampBits:   timestampBits,
		workerIdBits:    workerIdBits,
		sequenceBits:    sequenceBits,
		maxDeltaSeconds: -1 ^ (-1 << timestampBits),
		maxWorkerId:     -1 ^ (-1 << workerIdBits),
		maxSequence:     -1 ^ (-1 << sequenceBits),
		timestampShift:  workerIdBits + sequenceBits,
		workerIdShift:   sequenceBits,
	}

}

func (b *BitsAllocator) allocate(deltaSeconds, workerId, sequence int64) int64 {
	return (deltaSeconds << b.timestampShift) | (workerId << b.workerIdShift) | sequence
}
