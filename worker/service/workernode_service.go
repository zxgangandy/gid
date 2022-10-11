package service

import (
	"github.com/zxgangandy/gid/worker/dao"
	"github.com/zxgangandy/gid/worker/model"
	"gorm.io/gorm"
)

// IWorkerNodeService worker node service interface
type IWorkerNodeService interface {
	GetByHostname(HostName string) (*model.WorkerNode, error)
	Save(node *model.WorkerNode) (bool, error)
}

// WorkerNodeService worker node service
type WorkerNodeService struct {
	workerNodeDao dao.IWorkerNodeDao
}

// NewWorkerNodeService create worker node service instance
func NewWorkerNodeService(db *gorm.DB) *WorkerNodeService {
	workerNodeDao := dao.NewWorkerNodeDao(db)
	return &WorkerNodeService{
		workerNodeDao: workerNodeDao,
	}
}

// GetByHostname get worker node by hostname
func (w *WorkerNodeService) GetByHostname(HostName string) (*model.WorkerNode, error) {
	return w.workerNodeDao.GetByHostname(HostName)
}

// Save save  worker node
func (w *WorkerNodeService) Save(node *model.WorkerNode) (bool, error) {
	return w.workerNodeDao.Save(node)
}
