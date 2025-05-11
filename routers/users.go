package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	users := router.Group("/users")
	users.GET("/", api.ApiGroups.UsersAPI.UsersInfoView)
	users.POST("/register/", api.ApiGroups.UsersAPI.RegisterUser)
}
