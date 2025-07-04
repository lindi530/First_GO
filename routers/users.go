package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.POST("/register", api.ApiGroups.UserAPI.UserAccountAPI.Register)
	users.POST("/login", api.ApiGroups.UserAPI.UserAccountAPI.Login)

	users.Use(jwt.JWTAuthMiddleware())
	{
		users.POST("/logout", api.ApiGroups.UserAPI.UserAccountAPI.Logout)
		users.DELETE("/:user_id/posts/:post_id", api.ApiGroups.UserAPI.DeleteUserPost)
		users.DELETE("/:user_id", api.ApiGroups.UserAPI.UserAccountAPI.DeleteUser)
		users.GET("/:user_id", api.ApiGroups.UserAPI.UserInfo)
		users.GET("/userlist", api.ApiGroups.UserAPI.GetUserList)
		users.GET("/:user_id/posts", api.ApiGroups.UserAPI.GetUserPosts)
		users.POST("/:user_id/posts/create", api.ApiGroups.UserAPI.CreateUserPost)
		users.POST("/:user_id/modify_avatar", api.ApiGroups.UserAPI.ModifyAvatar)
		users.PATCH("/:user_id/profile", api.ApiGroups.UserAPI.ModifyProfile)

		users.POST("/online", api.ApiGroups.UserAPI.RefreshOnlineState)

		users.GET("/:user_id/is_following", api.ApiGroups.UserAPI.UserFollowAPI.CheckFollows)
		users.POST("/:user_id/follow", api.ApiGroups.UserAPI.UserFollowAPI.FollowUser)
		users.DELETE("/:user_id/follow", api.ApiGroups.UserAPI.UserFollowAPI.UnFollowUser)
	}

}
