package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func PostsRouter(router *gin.RouterGroup) {
	posts := router.Group("/posts")
	posts.GET("/:post_id/comments", api.ApiGroups.PostAPI.GetComments)
	posts.Use(jwt.JWTAuthMiddleware())
	{
		posts.GET("/page", api.ApiGroups.PostAPI.GetThePagePost)
		posts.GET("/:post_id", api.ApiGroups.PostAPI.GetOnePost)
		posts.GET("", api.ApiGroups.PostAPI.GetAllPost)
		posts.POST("/:post_id/comments", api.ApiGroups.PostAPI.CreateComment)
		posts.DELETE("/comments/:comment_id", api.ApiGroups.PostAPI.DeleteComment)
		posts.POST(":post_id/like", api.ApiGroups.PostAPI.PostLike)
		posts.POST(":post_id/unlike", api.ApiGroups.PostAPI.PostUnLike)
	}
}
