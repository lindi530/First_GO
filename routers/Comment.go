package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.RouterGroup) {
	comment := router.Group("/comment")

	comment.POST("/:comment_id/unlike", api.ApiGroups.CommentAPI.CommentUnLike)
	comment.POST("/:comment_id/like", api.ApiGroups.CommentAPI.CommentLike)
}
