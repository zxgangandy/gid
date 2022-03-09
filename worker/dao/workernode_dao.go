package dao

import (
	"github.com/zxgangandy/gid/worker/model"
	"gorm.io/gorm"
)

//worker node dao interface
type IWorkerNodeDao interface {
	GetByHostname(HostName string) (*model.WorkerNode, error)
	Save(node *model.WorkerNode) (bool, error)
}

//worker node dao implementation
type WorkerNodeDao struct {
	db *gorm.DB
}

//create a worker node instance
func NewWorkerNodeDao(db *gorm.DB) *WorkerNodeDao {
	return &WorkerNodeDao{db: db}
}

//get worker node by host name
func (w *WorkerNodeDao) GetByHostname(HostName string) (*model.WorkerNode, error) {
	var node model.WorkerNode
	err := w.db.Where("host_name =? ", HostName).Find(&node).Error
	return &node, err
}

//save worker node
func (w *WorkerNodeDao) Save(node *model.WorkerNode) (bool, error) {
	err := w.db.Save(node).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
