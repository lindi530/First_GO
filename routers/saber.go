package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func SaberRouter(router *gin.RouterGroup) {
	saber := router.Group("/saber")

	saber.Use(jwt.JWTAuthMiddleware())
	{
		saber.GET("info", api.ApiGroups.SaberAPI.GetSaberStat)
	}
}
