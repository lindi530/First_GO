package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func PostsRouter(router *gin.RouterGroup) {
	posts := router.Group("/posts")
	posts.GET("", api.ApiGroups.PostAPI.GetAllPost)
	posts.GET("/page", api.ApiGroups.PostAPI.GetThePagePost)
}
