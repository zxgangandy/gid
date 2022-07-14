package tests

import (
	"fmt"
	"github.com/zxgangandy/gid"
	"github.com/zxgangandy/gid/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func TestDefaultGenId(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/jingwei-exchange?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	c := config.New(db, "8000")

	start := time.Now().UnixMilli()
	gen := gid.New(c)
	for i := 0; i < 10000; i++ {
		_ = gen.GetUID()
	}
	end := time.Now().UnixMilli()
	fmt.Println(end - start)
}

func TestCustomGenId(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/jingwei-exchange?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	c := config.New(db, "8000")
	c.WorkerBits = 10
	c.SeqBits = 23

	start := time.Now().UnixMilli()
	gen := gid.New(c)
	for i := 0; i < 10000000; i++ {
		_ = gen.GetUID()
	}
	end := time.Now().UnixMilli()
	fmt.Println(end - start)
}
