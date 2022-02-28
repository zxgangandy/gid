package worker

import (
	"gid"
	"gid/worker/model"
	"gid/worker/service"
	"time"
)

type IdAssigner interface {
	AssignWorkerId() int64
}

type DisposableWorkerIdAssigner struct {
	workerNodeService service.IWorkerNodeService
	config            gid.Config
}

func NewWorkerIdAssigner(config gid.Config) *DisposableWorkerIdAssigner {
	workerNodeService := service.NewWorkerNodeService(config.GetDB())
	return &DisposableWorkerIdAssigner{
		workerNodeService: workerNodeService,
		config:            config,
	}
}

func (d *DisposableWorkerIdAssigner) AssignWorkerId() int64 {
	node, err := d.workerNodeService.GetByHostname(d.config.GetHostName())
	if err != nil {
		panic(err)
	}

	if node != nil {
		return node.Id
	}

	newNode := d.buildWorkerNode(d.config)
	_, saveErr := d.workerNodeService.Save(newNode)
	if saveErr != nil {
		panic(saveErr)
	}

	return newNode.Id
}

func (d *DisposableWorkerIdAssigner) buildWorkerNode(config gid.Config) *model.WorkerNode {
	node := &model.WorkerNode{
		Type:       ACTUAL,
		LaunchDate: time.Now(),
	}

	if gid.IsDocker {
		node.HostName = gid.DockerHost
		node.Port = gid.DockerPort
		node.Type = CONTAINER
	} else {
		node.Type = ACTUAL
		node.HostName = gid.HostName
		node.Port = config.GetPort()
	}

	return node
}
