package main

import (
	"fmt"
	"github.com/catbugdemo/project_order/inits"
)

func main() {

	go inits.InitHttp()

	fmt.Println("hello world")

	select {}
}

