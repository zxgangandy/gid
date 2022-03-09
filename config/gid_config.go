package config

import "gorm.io/gorm"

// uid configure interface
type UidConfig interface {
	//get db handler
	GetDB() *gorm.DB

	//get port
	GetPort() string

	//get time bits
	GetTimeBits() uint32

	//get worker bits
	GetWorkerBits() uint32

	//get sequence bits
	GetSeqBits() uint32

	//get seconds of the epoch time
	GetEpochSeconds() int64

	//get max backward seconds
	GetMaxBackwardSeconds() int64

	//get enable backward status
	EnableBackward() bool
}

// the default uid configure
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

//create a default uid configure instance
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

//get db handler
func (c *DefaultUidConfig) GetDB() *gorm.DB {
	return c.DB
}

//get app port
func (c *DefaultUidConfig) GetPort() string {
	return c.Port
}

//get time bits
func (c *DefaultUidConfig) GetTimeBits() uint32 {
	return c.TimeBits
}

//get worker bits
func (c *DefaultUidConfig) GetWorkerBits() uint32 {
	return c.WorkerBits
}

//get sequence bits
func (c *DefaultUidConfig) GetSeqBits() uint32 {
	return c.SeqBits
}

//get seconds of the epoch time
func (c *DefaultUidConfig) GetEpochSeconds() int64 {
	return c.EpochSeconds
}

//get the seconds of the max backward
func (c *DefaultUidConfig) GetMaxBackwardSeconds() int64 {
	return c.MaxBackwardSeconds
}

//get enable backward status
func (c *DefaultUidConfig) EnableBackward() bool {
	return c.IsEnableBackward
}
