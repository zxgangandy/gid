package dao

import (
	"gid/worker/model"
	"gorm.io/gorm"
)

type IWorkerNodeDao interface {
	GetByHostname(HostName string) (*model.WorkerNode, error)
	Save(node *model.WorkerNode) (bool, error)
}

type WorkerNodeDao struct {
	db *gorm.DB
}

func NewWorkerNodeDao(db *gorm.DB) *WorkerNodeDao {
	return &WorkerNodeDao{db: db}
}

func (w *WorkerNodeDao) GetByHostname(HostName string) (*model.WorkerNode, error) {
	var node *model.WorkerNode
	err := w.db.Where("host_name =? ", HostName).Find(node).Error
	return node, err
}

func (w *WorkerNodeDao) Save(node *model.WorkerNode) (bool, error) {
	err := w.db.Save(node).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
