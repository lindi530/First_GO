package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	users := router.Group("/users")
	users.GET("/", api.ApiGroups.UserAPI.UserInfo)
	users.POST("/register/", api.ApiGroups.UserAPI.RegisterUser)
	users.POST("/login/", api.ApiGroups.UserAPI.LoginUser)
	users.DELETE("/:id", api.ApiGroups.UserAPI.DeleteUser)
}
