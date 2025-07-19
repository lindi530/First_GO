package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func FollowRouter(router *gin.RouterGroup) {
	follow := router.Group("/follow")

	follow.Use(jwt.JWTAuthMiddleware())
	{
		follow.GET("userlist", api.ApiGroups.UserAPI.UserFollowAPI.GetFollowUsers)
	}
}
