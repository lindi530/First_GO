package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func MatchRouter(router *gin.RouterGroup) {
	match := router.Group("/match")
	match.POST("", api.ApiGroups.MatchAPI.GetMatch)
}
