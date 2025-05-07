package main

import (
	"GO1/core"
	"GO1/global"
	"GO1/routers"
	"GO1/test"
	"fmt"
)

func main() {
	Start()
	router := routers.InitRouter()
	api := router.Group("/")
	test.Test(api)

	addr := global.Config.System.Addr()
	global.Logger.Info(fmt.Sprintf("Gin 运行在：%s", addr))
	router.Run(addr)

}

func Start() {
	core.InitConfig()
	core.InitGorm()
	core.InitLogger()
}
