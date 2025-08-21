package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func SaberRouter(router *gin.RouterGroup) {
	saber := router.Group("/saber")
	rooms := saber.Group("/rooms")

	saber.Use(jwt.JWTAuthMiddleware())
	{
		saber.GET("/info", api.ApiGroups.SaberAPI.GetSaberStat)
		saber.POST("/submit", api.ApiGroups.SaberAPI.SaberSubmit)
		saber.POST("/invite", api.ApiGroups.SaberAPI.InviteFriend)
	}

	rooms.Use(jwt.JWTAuthMiddleware())
	{
		rooms.POST("/expire/:room_id", api.ApiGroups.SaberAPI.ExpireRoom)
	}
}
