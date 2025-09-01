package comment_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/comment_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentAPI) CommentUnLike(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	commentId, _ := strconv.ParseInt(c.Param("comment_id"), 10, 64)

	result := comment_service.CommentUnLike(userId, commentId)
	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.Ok(result.Data, result.Msg, c)
}
