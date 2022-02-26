package gid

import "gorm.io/gorm"

type Config interface {
	GetHostName() string

	GetDB() *gorm.DB

	GetPort() int32
}
