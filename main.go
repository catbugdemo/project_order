package main

import (
	"fmt"
	"runtime"

	"github.com/catbugdemo/project_order/inits"
)

// @title 初始项目
// @version 1.0
// @description 声明（可不写）
// @termsOfService

// @contact.name www.github.com
// @contact.url https://www.github.com
// @contact.email me@test.me

// @license.name Apache 2.0
// @license.url

// @host 127.0.0.1:8080
// @BasePath

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go inits.InitGin()

	fmt.Println("hello world")

	select {}
}
