package config

import "gorm.io/gorm"

type UidConfig interface {
	GetDB() *gorm.DB

	GetPort() string

	GetTimeBits() uint32

	GetWorkerBits() uint32

	GetSeqBits() uint32

	GetEpochSeconds() int64

	GetMaxBackwardSeconds() int64

	EnableBackward() bool
}
