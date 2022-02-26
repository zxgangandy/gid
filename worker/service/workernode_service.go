package service

import (
	"gid/worker/dao"
	"gid/worker/model"
	"gorm.io/gorm"
)

type IWorkerNodeService interface {
	GetByHostname(HostName string) (*model.WorkerNode, error)
	Save(node *model.WorkerNode) (bool, error)
}

type WorkerNodeService struct {
	workerNodeDao dao.IWorkerNodeDao
}

func NewWorkerNodeService(db *gorm.DB) *WorkerNodeService {
	workerNodeDao := dao.NewWorkerNodeDao(db)
	return &WorkerNodeService{
		workerNodeDao: workerNodeDao,
	}
}

func (w *WorkerNodeService) GetByHostname(HostName string) (*model.WorkerNode, error) {
	return w.workerNodeDao.GetByHostname(HostName)
}

func (w *WorkerNodeService) Save(node *model.WorkerNode) (bool, error) {
	return w.workerNodeDao.Save(node)
}
