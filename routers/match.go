package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func MatchRouter(router *gin.RouterGroup) {
	match := router.Group("/match")
	match.Use(jwt.JWTAuthMiddleware())
	{
		match.POST("", api.ApiGroups.MatchAPI.MatchRequest)
	}

}
