package gid

import (
	"errors"
	"fmt"
	"github.com/zxgangandy/gid/config"
	"github.com/zxgangandy/gid/worker"
	"sync"
	"time"
)

// UidGenerator generator interface
type UidGenerator interface {
	GetUID() int64

	ParseUID(uid int64) string
}

// DefaultUidGenerator default uid generator
type DefaultUidGenerator struct {
	bitsAllocator *BitsAllocator
	config        config.UidConfig
	mutex         sync.Mutex
	workerId      int64
	lastSecond    int64
	sequence      int64
}

// New create the default uid generator instance
func New(config config.UidConfig) *DefaultUidGenerator {
	idAssigner := worker.NewWorkerIdAssigner(config)
	allocator := NewBitsAllocator(config.GetTimeBits(), config.GetWorkerBits(), config.GetSeqBits())

	var workerId int64
	workerId = idAssigner.AssignWorkerId()
	if workerId > allocator.maxWorkerId {
		workerId = workerId % allocator.maxWorkerId
	}

	return &DefaultUidGenerator{
		bitsAllocator: allocator,
		config:        config,
		workerId:      workerId,
		sequence:      0,
	}
}

// GetUID generate the unique id
func (g *DefaultUidGenerator) GetUID() int64 {
	c := g.config
	return g.nextId(c.GetEpochSeconds(), c.GetMaxBackwardSeconds(), c.EnableBackward())
}

// ParseUID parse the generated unique id then get the meta information
// +------+----------------------+----------------+-----------+
// | sign |     delta seconds    | worker node id | sequence  |
// +------+----------------------+----------------+-----------+
//   1bit          30bits              7bits         13bits
func (g *DefaultUidGenerator) ParseUID(uid int64) string {
	totalBits := (uint32)(TotalBits)
	allocateTotalBits := g.bitsAllocator.allocateTotalBits
	signBits := g.bitsAllocator.signBits
	timestampBits := g.bitsAllocator.timestampBits
	workerIdBits := g.bitsAllocator.workerIdBits
	sequenceBits := g.bitsAllocator.sequenceBits

	sequence := uint64(uid<<(totalBits-sequenceBits)) >> (totalBits - sequenceBits)
	workerId := uint64(uid<<(timestampBits+signBits+totalBits-allocateTotalBits)) >> (totalBits - workerIdBits)
	deltaSeconds := int64(uint64(uid) >> (workerIdBits + sequenceBits))

	// format as string
	return fmt.Sprintf(
		"{\"uid\":\"%d\",\"timestamp\":\"%d\",\"workerId\":\"%d\",\"sequence\":\"%d\"}",
		uid, g.config.GetEpochSeconds()+deltaSeconds, workerId, sequence,
	)
}

func (g *DefaultUidGenerator) nextId(epochSeconds, maxBackwardSeconds int64, enableBackward bool) int64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	currentSecond := g.getCurrentSecond(epochSeconds)

	if currentSecond < g.lastSecond {
		refusedSeconds := g.lastSecond - currentSecond
		if !enableBackward {
			panic(errors.New(fmt.Sprintf("Clock moved backwards. Refusing for %d seconds", refusedSeconds)))
		}

		if refusedSeconds <= maxBackwardSeconds {
			for currentSecond < g.lastSecond {
				currentSecond = g.getCurrentSecond(epochSeconds)
			}
		} else {
			panic("Clock moved backwards. Refused seconds bigger than max backward seconds")
		}
	} else if currentSecond == g.lastSecond { // At the same second, increase sequence
		g.sequence = (g.sequence + 1) & g.bitsAllocator.maxSequence
		// Exceed the max sequence, we wait the next second to generate uid
		if g.sequence == 0 {
			currentSecond = g.getNextSecond(g.lastSecond, epochSeconds)
		}
	} else {
		// At the different second, sequence restart from zero
		g.sequence = 0
	}

	g.lastSecond = currentSecond

	// Allocate bits for UID
	return g.bitsAllocator.allocate(currentSecond-epochSeconds, g.workerId, g.sequence)
}

func (g *DefaultUidGenerator) getCurrentSecond(epochSeconds int64) int64 {
	currentSeconds := time.Now().Unix()
	if currentSeconds-epochSeconds > g.bitsAllocator.maxDeltaSeconds {
		panic("Timestamp bits is exhausted. Refusing UID generate.")
	}

	return currentSeconds
}

func (g *DefaultUidGenerator) getNextSecond(lastTimestamp, epochSeconds int64) int64 {
	timestamp := g.getCurrentSecond(epochSeconds)
	for timestamp <= lastTimestamp {
		timestamp = g.getCurrentSecond(epochSeconds)
	}

	return timestamp
}
