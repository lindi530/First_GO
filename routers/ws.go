package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func WSRouter(router *gin.RouterGroup) {
	chat := router.Group("/ws")
	chat.GET("", api.ApiGroups.ChatAPI.ChatHandler)
	chat.POST("reconnect", api.ApiGroups.ChatAPI.ReconnectHandler)
}
