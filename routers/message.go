package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func MessageRouter(router *gin.RouterGroup) {
	message := router.Group("/message")
	message.Use(jwt.JWTAuthMiddleware())
	{
		message.GET("/:target_id", api.ApiGroups.MessageAPI.GetMessageByTargetId)
	}
}
