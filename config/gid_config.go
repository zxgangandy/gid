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
	DB                 *gorm.DB
	Port               string // App port
	TimeBits           uint32 // time bits
	WorkerBits         uint32 // worker bits
	SeqBits            uint32 // sequence bits
	EpochSeconds       int64  // epoch seconds
	MaxBackwardSeconds int64  // max backward seconds
	IsEnableBackward   bool   // enable clock backward
}

func New(db *gorm.DB, port string) *DefaultUidConfig {
	return &DefaultUidConfig{
		DB:                 db,
		Port:               port,
		TimeBits:           30,
		WorkerBits:         7,
		SeqBits:            13,
		EpochSeconds:       1550592000000 / 1000,
		MaxBackwardSeconds: 1,
		IsEnableBackward:   true,
	}
}

func (c *DefaultUidConfig) GetDB() *gorm.DB {
	return c.DB
}

func (c *DefaultUidConfig) GetPort() string {
	return c.Port
}

func (c *DefaultUidConfig) GetTimeBits() uint32 {
	return c.TimeBits
}

func (c *DefaultUidConfig) GetWorkerBits() uint32 {
	return c.WorkerBits
}

func (c *DefaultUidConfig) GetSeqBits() uint32 {
	return c.SeqBits
}

func (c *DefaultUidConfig) GetEpochSeconds() int64 {
	return c.EpochSeconds
}

func (c *DefaultUidConfig) GetMaxBackwardSeconds() int64 {
	return c.MaxBackwardSeconds
}

func (c *DefaultUidConfig) EnableBackward() bool {
	return c.IsEnableBackward
}
