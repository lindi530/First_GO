package main

import (
	"GO1/global"
	"GO1/pkg/conf"
	"GO1/pkg/gorm"
	"GO1/pkg/logger"
	"GO1/pkg/redis"
	"GO1/pkg/snowflake"
	"GO1/pkg/translator"
	"GO1/pkg/validator"
	"GO1/routers"
	"GO1/test"
)

func main() {
	Start()
	Define()
	router := routers.InitRouter()
	api := router.Group("/")

	if global.Config.System.Env == "test" {
		test.Test(api)
	}
	addr := global.Config.System.Addr()
	//global.Logger.Info(fmt.Sprintf("Gin 运行在：%s", addr))
	router.Run(addr)
}

func Start() {
	conf.InitConfig()
	gorm.InitGorm()
	logger.InitLogger()
	snowflake.InitSnowFlake()
	validator.InitValidator()
	translator.InitTrans("zh")
	redis.InitRedisClient()
}

func Define() {
	validator.DefinedValidator()
}
