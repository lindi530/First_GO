package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	api := router.Group("/")
	AddRouter(api)
	return router
}

func AddRouter(api *gin.RouterGroup) {
	SettingsRouter(api)
	UsersRouter(api)
}
