package post_api

import (
	"GO1/middlewares/response"
	service_post "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) GetComments(c *gin.Context) {
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	result := service_post.GetComments(postId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithData(result.Data, c)
}
