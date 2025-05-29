package post_api

import (
	"GO1/middlewares/response"
	service_post "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) DeleteComment(c *gin.Context) {
	commentId, _ := strconv.ParseInt(c.Param("comment_id"), 10, 64)

	// 服务
	result := service_post.DeleteComment(c, commentId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithMessage(result.Msg, c)
}
