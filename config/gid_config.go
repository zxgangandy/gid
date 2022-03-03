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

type DefaultUidConfig struct {
	db                 *gorm.DB
	port               string
	timeBits           uint32
	workerBits         uint32
	seqBits            uint32
	epochSeconds       int64
	maxBackwardSeconds int64
	enableBackward     bool
}

func New(db *gorm.DB, port string) *DefaultUidConfig {
	return &DefaultUidConfig{
		db:                 db,
		port:               port,
		timeBits:           30,
		workerBits:         10,
		seqBits:            13,
		epochSeconds:       1550592000000 / 1000,
		maxBackwardSeconds: 1,
		enableBackward:     true,
	}
}

func (c *DefaultUidConfig) GetDB() *gorm.DB {
	return c.db
}

func (c *DefaultUidConfig) GetPort() string {
	return c.port
}

func (c *DefaultUidConfig) GetTimeBits() uint32 {
	return c.timeBits
}

func (c *DefaultUidConfig) GetWorkerBits() uint32 {
	return c.workerBits
}

func (c *DefaultUidConfig) GetSeqBits() uint32 {
	return c.seqBits
}

func (c *DefaultUidConfig) GetEpochSeconds() int64 {
	return c.epochSeconds
}

func (c *DefaultUidConfig) GetMaxBackwardSeconds() int64 {
	return c.maxBackwardSeconds
}

func (c *DefaultUidConfig) EnableBackward() bool {
	return c.enableBackward
}
