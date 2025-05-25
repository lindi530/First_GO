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

	users.Use(jwt.JWTAuthMiddleware())
	{
		users.DELETE("/:user_id/posts/:post_id", api.ApiGroups.UserAPI.DeleteUserPost)
		users.DELETE("/:user_id", api.ApiGroups.UserAPI.DeleteUser)
		users.GET("/:user_id", api.ApiGroups.UserAPI.UserInfo)
		users.GET("/userlist", api.ApiGroups.UserAPI.GetUserList)
		users.GET("/:user_id/posts", api.ApiGroups.UserAPI.GetUserPosts)
		users.POST("/:user_id/posts/create", api.ApiGroups.UserAPI.CreateUserPost)
		users.POST("/:user_id/modify_avatar", api.ApiGroups.UserAPI.ModifyAvatar)
		users.PATCH("/:user_id/profile", api.ApiGroups.UserAPI.ModifyProfile)
	}

}
