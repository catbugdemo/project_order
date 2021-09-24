package main

import (
	"fmt"
	"runtime"

	"github.com/catbugdemo/project_order/inits"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go inits.InitGin()

	fmt.Println("hello world")

	select {}
}
