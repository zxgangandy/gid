package worker

import (
	"github.com/zxgangandy/gid/config"
	"github.com/zxgangandy/gid/worker/model"
	"github.com/zxgangandy/gid/worker/service"
	"time"
)

type IdAssigner interface {
	AssignWorkerId() int64
}

type DisposableWorkerIdAssigner struct {
	workerNodeService service.IWorkerNodeService
	config            config.UidConfig
}

func NewWorkerIdAssigner(config config.UidConfig) *DisposableWorkerIdAssigner {
	workerNodeService := service.NewWorkerNodeService(config.GetDB())
	return &DisposableWorkerIdAssigner{
		workerNodeService: workerNodeService,
		config:            config,
	}
}

func (d *DisposableWorkerIdAssigner) AssignWorkerId() int64 {
	newNode := d.buildWorkerNode(d.config.GetPort())
	node, err := d.workerNodeService.GetByHostname(newNode.HostName)
	if err != nil {
		panic(err)
	}

	if node != nil {
		return node.Id
	}

	_, saveErr := d.workerNodeService.Save(newNode)
	if saveErr != nil {
		panic(saveErr)
	}

	return newNode.Id
}

func (d *DisposableWorkerIdAssigner) buildWorkerNode(port string) *model.WorkerNode {
	node := &model.WorkerNode{
		Type:       ACTUAL,
		LaunchDate: time.Now(),
	}

	if config.IsDocker {
		node.HostName = config.DockerHost
		node.Port = config.DockerPort
		node.Type = CONTAINER
	} else {
		node.Type = ACTUAL
		node.HostName = config.HostName
		node.Port = port
	}

	return node
}
