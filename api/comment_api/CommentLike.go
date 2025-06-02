package comment_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	service_comment "GO1/service/Comment"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentAPI) CommentLike(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	commentId, _ := strconv.ParseInt(c.Param("comment_id"), 10, 64)

	result := service_comment.CommentLike(userId, commentId)
	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.Ok(result.Data, result.Msg, c)
}
