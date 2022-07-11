package config

import "gorm.io/gorm"

// UidConfig uid configure interface
type UidConfig interface {
	// GetDB get db handler
	GetDB() *gorm.DB

	// GetPort get port
	GetPort() string

	// GetTimeBits get time bits
	GetTimeBits() uint32

	// GetWorkerBits get worker bits
	GetWorkerBits() uint32

	// GetSeqBits get sequence bits
	GetSeqBits() uint32

	// GetEpochSeconds get seconds of the epoch time
	GetEpochSeconds() int64

	// GetMaxBackwardSeconds get max backward seconds
	GetMaxBackwardSeconds() int64

	// EnableBackward get enable backward status
	EnableBackward() bool
}

// DefaultUidConfig the default uid configure
type DefaultUidConfig struct {
	DB                 *gorm.DB // db handler
	Port               string   // app port
	TimeBits           uint32   // time bits
	WorkerBits         uint32   // worker bits
	SeqBits            uint32   // sequence bits
	EpochSeconds       int64    // epoch seconds
	MaxBackwardSeconds int64    // max backward seconds
	IsEnableBackward   bool     // enable clock backward
}

// New create a default uid configure instance
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

// GetDB get db handler
func (c *DefaultUidConfig) GetDB() *gorm.DB {
	return c.DB
}

// GetPort get app port
func (c *DefaultUidConfig) GetPort() string {
	return c.Port
}

// GetTimeBits get time bits
func (c *DefaultUidConfig) GetTimeBits() uint32 {
	return c.TimeBits
}

// GetWorkerBits get worker bits
func (c *DefaultUidConfig) GetWorkerBits() uint32 {
	return c.WorkerBits
}

// GetSeqBits get sequence bits
func (c *DefaultUidConfig) GetSeqBits() uint32 {
	return c.SeqBits
}

// GetEpochSeconds get seconds of the epoch time
func (c *DefaultUidConfig) GetEpochSeconds() int64 {
	return c.EpochSeconds
}

// GetMaxBackwardSeconds get the seconds of the max backward
func (c *DefaultUidConfig) GetMaxBackwardSeconds() int64 {
	return c.MaxBackwardSeconds
}

// EnableBackward get enable backward status
func (c *DefaultUidConfig) EnableBackward() bool {
	return c.IsEnableBackward
}
