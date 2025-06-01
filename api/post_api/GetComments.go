package post_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	service_post "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) GetComments(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	result := service_post.GetComments(c, userId, postId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithData(result.Data, c)
}
