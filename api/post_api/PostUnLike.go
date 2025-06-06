package post_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) PostUnLike(c *gin.Context) {
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := post.PostUnLike(userId, postId)

	if !resp.Ok {
		response.FailWithMessage(resp.Msg, c)
		return
	}
	response.OkWithMessage(resp.Msg, c)
}
