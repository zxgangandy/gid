package model

import (
	"time"
)

type WorkerNode struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	HostName   string    `gorm:"column:host_name"`
	Port       string    `gorm:"column:port"`
	Type       int32     `gorm:"column:type"`
	LaunchDate time.Time `gorm:"column:launch_date"`
	Modified   time.Time `gorm:"column:modified"`
	Created    time.Time `gorm:"column:created"`
}
