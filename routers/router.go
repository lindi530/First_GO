package routers

import (
	"GO1/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)

	router := gin.New()
	api := router.Group("/")
	AddRouter(api)

	return router
}

func AddRouter(api *gin.RouterGroup) {
	SettingsRouter(api)
	UsersRouter(api)
}
