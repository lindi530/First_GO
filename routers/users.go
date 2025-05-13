package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.POST("/register/", api.ApiGroups.UserAPI.RegisterUser)
	users.POST("/login/", api.ApiGroups.UserAPI.LoginUser)
	users.DELETE("/:id", api.ApiGroups.UserAPI.DeleteUser)

	users.Use(jwt.JWTAuthMiddleware())
	{
		users.GET("/:id", api.ApiGroups.UserAPI.UserInfo)
	}

}
