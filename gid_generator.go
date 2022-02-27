package gid

import (
	"gid/worker"
	"gid/worker/service"
)

type UidGenerator interface {
	GetUID() (uint64, error)

	ParseUID(uid uint64) (string, error)
}

type DefaultUidGenerator struct {
	workerNodeService service.IWorkerNodeService
}

func NewUidGenerator(config Config) *DefaultUidGenerator {

	idAssigner := worker.NewWorkerIdAssigner(config)

	allocator := NewBitsAllocator(config.GetTimeBits(), config.GetWorkerBits(), config.GetSeqBits())
	workerId := idAssigner.AssignWorkerId()

	if workerId > allocator.MaxWorkerId {
		workerId = workerId % allocator.MaxWorkerId
	}

	workerNodeService := service.NewWorkerNodeService(config.GetDB())
	return &DefaultUidGenerator{
		workerNodeService: workerNodeService,
	}
}

func (g *DefaultUidGenerator) GetUID() (uint64, error) {

	return 0, nil
}

func (g *DefaultUidGenerator) ParseUID() (string, error) {
	return "", nil
}

func (g *DefaultUidGenerator) InitGenerator(config Config) error {

	return nil
}
