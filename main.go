package main

import (
	"GO1/core"
	"GO1/middlewares"
	"GO1/test"
	"github.com/gin-gonic/gin"
)

func main() {
	Start()
	router := gin.New()
	test.Test(router.Group("/"))
	router.Use(gin.Recovery(), middlewares.Logger())

	router.Run(":8000")

}

func Start() {
	core.InitConfig()
	core.InitGorm()
	core.InitLogger()
}
