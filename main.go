package main

import (
	"GO1/database/sync"
	"GO1/global"
	"GO1/pkg/conf"
	"GO1/pkg/gorm"
	"GO1/pkg/jwt"
	"GO1/pkg/logger"
	"GO1/pkg/rabbitmq"
	"GO1/pkg/redis"
	"GO1/pkg/snowflake"
	"GO1/pkg/translator"
	"GO1/pkg/validator"
	"GO1/routers"
	"GO1/service/match_service"
	"GO1/service/ws_service"
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	initDependencies()        //初始化依赖
	initCustomValidator()     // 注册自定义验证器
	go ws_service.WsHub.Run() // go 协程
	go startCronJob()         // 定时作业
	go syncInitialData()      // 初始同步数据
	startServer()
}

func initDependencies() {
	conf.InitConfig()
	logger.InitLogger()
	gorm.InitGorm()
	snowflake.InitSnowFlake()
	validator.InitValidator()
	translator.InitTrans("zh")
	redis.InitRedisClient()
	jwt.InitJwt()
	rabbitmq.InitRabbitMQ()
}

func initCustomValidator() {
	validator.DefinedValidator()
}

func startCronJob() {
	c := cron.New()
	// 每天凌晨 3 点执行
	c.AddFunc("0 3 * * *", sync.SyncCommentLikes)
	c.Start()
}

func syncInitialData() {
	sync.SyncCommentLikes()
	sync.SyncPostLikes()
}

func startServer() {
	match_service.StartMatchConsumer()
	startHTTPServer() // 必须在最后
}

func startHTTPServer() {
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Logger.Info(fmt.Sprintf("Gin 运行在：%s", addr))
	if err := router.Run(addr); err != nil {
		global.Logger.Fatalf("服务器启动失败: %v", err)
	}
}
