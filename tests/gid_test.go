package tests

import (
	"fmt"
	"github.com/zxgangandy/gid"
	"github.com/zxgangandy/gid/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestGenId(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/jingwei-exchange?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	c := config.New(db, "8000")
	fmt.Println(gid.New(c).GetUID())

	c.WorkerBits = 20
	c.SeqBits = 13

	fmt.Println(gid.New(c).GetUID())
}
