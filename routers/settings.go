package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func SettingsRouter(router *gin.RouterGroup) {
	settings := router.Group("/settings")
	settings.GET("/", api.ApiGroups.SettingsAPI.SettingsInfoView)
}
