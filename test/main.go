package main

import (
	"fmt"
	"gid"
	"gorm.io/gorm"
)

type TestConfig struct {
}

func newTestConfig() *TestConfig {
	return &TestConfig{}
}

func (c *TestConfig) GetDB() *gorm.DB {
	return nil //add your real db handle
}

func (c *TestConfig) GetPort() string {
	return "8000"
}

func (c *TestConfig) GetTimeBits() uint32 {
	return 30
}

func (c *TestConfig) GetWorkerBits() uint32 {
	return 16
}

func (c *TestConfig) GetSeqBits() uint32 {
	return 7
}

func (c *TestConfig) GetEpochSeconds() int64 {
	return 1550592000000 / 1000
}

func (c *TestConfig) GetMaxBackwardSeconds() int64 {
	return 1
}

func (c *TestConfig) EnableBackward() bool {
	return true
}

func main() {
	config := newTestConfig()
	gid.New(config)
	fmt.Println("main")
}
