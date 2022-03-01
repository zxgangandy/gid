package main

import (
	"fmt"
	"github.com/zxgangandy/gid"
	"github.com/zxgangandy/gid/config"
)

func main() {
	c := config.New(nil, "8000")
	gid.New(c).GetUID()
	fmt.Println("main")
}
