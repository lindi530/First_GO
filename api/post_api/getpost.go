package post_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/post"
	"GO1/pkg/jwt"
	service "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) GetAllPost(c *gin.Context) {
	// 数据校验
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	// 服务
	posts, _ := service.GetAllPost(c, userId)

	// 响应
	response.OkWithData(posts, c)
}

func (PostAPI) GetOnePost(c *gin.Context) {
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	result := service.GetOnePost(c, userId, postId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithData(result.Data, c)
}

func (PostAPI) GetThePagePost(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	var pageInfo post.PageInfo
	if err := c.ShouldBind(&pageInfo); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}
	global.Logger.Infof("GetThePagePost pageInfo: %v", pageInfo)
	posts, err := service.GetThePagePost(c, userId, pageInfo)

	if err != nil {
		response.FailWithMessage("数据拉取失败", c)
		return
	}

	response.OkWithData(posts, c)
}
