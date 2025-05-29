package post_api

import (
	"GO1/middlewares/response"
	"GO1/models/Comment"
	"GO1/pkg/jwt"
	service_user "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) CreateComment(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	var comment Comment.Comment
	if c.ShouldBindJSON(&comment) != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	result := service_user.CreateComment(userId, postId, &comment)
	if result.Ok == false {
		response.FailWithMessage(result.Msg, c)
		return
	}

	response.OkWithData(comment, c)
}
