package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.POST("/register", api.ApiGroups.UserAPI.Register)
	users.POST("/login", api.ApiGroups.UserAPI.Login)
	users.POST("/refresh", api.ApiGroups.UserAPI.Refresh)
	users.POST("/logout", api.ApiGroups.UserAPI.Logout)
	users.DELETE("/:id", api.ApiGroups.UserAPI.Delete)

	users.Use(jwt.JWTAuthMiddleware())
	{
		users.GET("/:id", api.ApiGroups.UserAPI.UserInfo)
		users.GET("/:id/posts", api.ApiGroups.UserAPI.GetUserPosts)
		users.POST("/:id/post", api.ApiGroups.UserAPI.SaveUserPost)
	}

}
