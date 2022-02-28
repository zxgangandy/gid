package gid

import "gorm.io/gorm"

type Config interface {
	GetHostName() string

	GetDB() *gorm.DB

	GetPort() string

	GetTimeBits() uint32

	GetWorkerBits() uint32

	GetSeqBits() uint32

	GetEpochSeconds() int64

	GetMaxBackwardSeconds() int64

	EnableBackward() bool
}
