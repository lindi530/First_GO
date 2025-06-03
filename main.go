package main

import (
	"GO1/database/sync"
	"GO1/global"
	"GO1/pkg/conf"
	"GO1/pkg/gorm"
	"GO1/pkg/logger"
	"GO1/pkg/redis"
	"GO1/pkg/snowflake"
	"GO1/pkg/translator"
	"GO1/pkg/validator"
	"GO1/routers"
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	Start()
	Define()

	router := routers.InitRouter()
	sync.SyncCommentLikes()
	go StartCronJob()

	//if global.Config.System.Env == "test" {
	//	test.Test(api)
	//}
	addr := global.Config.System.Addr()

	global.Logger.Info(fmt.Sprintf("Gin 运行在：%s", addr))
	router.Run(addr)
}

func Start() {
	conf.InitConfig()
	logger.InitLogger()
	gorm.InitGorm()
	snowflake.InitSnowFlake()
	validator.InitValidator()
	translator.InitTrans("zh")
	redis.InitRedisClient()
}

func Define() {
	validator.DefinedValidator()
}

func StartCronJob() {
	c := cron.New()
	// 每天凌晨 3 点执行
	c.AddFunc("0 3 * * *", sync.SyncCommentLikes)
	c.Start()
}
