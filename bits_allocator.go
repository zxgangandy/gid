package gid

import "errors"

const TotalBits = 1 << 6

type BitsAllocator struct {
	TimestampBits uint32
	WorkerIdBits  uint32
	SequenceBits  uint32

	/**
	 * Max value for workId & sequence
	 */
	MaxDeltaSeconds uint32
	MaxWorkerId     uint64
	MaxSequence     uint32

	/**
	 * Shift for timestamp & workerId
	 */
	TimestampShift uint32
	WorkerIdShift  uint32
}

func NewBitsAllocator(timestampBits, workerIdBits, sequenceBits uint32) *BitsAllocator {
	var signBits uint32 = 1

	allocateTotalBits := signBits + timestampBits + workerIdBits + sequenceBits

	if allocateTotalBits > TotalBits {
		panic(errors.New("allocate larger than 64 bits"))
	}

	return &BitsAllocator{
		TimestampBits:   timestampBits,
		WorkerIdBits:    workerIdBits,
		SequenceBits:    sequenceBits,
		MaxDeltaSeconds: ^(-1 << timestampBits),
		MaxWorkerId:     ^(-1 << workerIdBits),
		MaxSequence:     ^(-1 << sequenceBits),
		TimestampShift:  workerIdBits + sequenceBits,
		WorkerIdShift:   sequenceBits,
	}

}
