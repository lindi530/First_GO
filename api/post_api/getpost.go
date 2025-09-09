package post_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/post_model"
	"GO1/pkg/jwt"
	service "GO1/service/user_service/user_post_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) GetAllPost(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	// 服务
	posts, _ := service.GetAllPost(userId)

	// 响应
	response.OkWithData(posts, c)
}

func (PostAPI) GetOnePost(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)
	result := service.GetOnePost(userId, postId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithData(result.Data, c)
}

func (PostAPI) GetThePagePost(c *gin.Context) {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	var pageInfo post_model.PageInfo
	if err := c.ShouldBind(&pageInfo); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}
	global.Logger.Infof("GetThePagePost pageInfo: %v", pageInfo)
	posts, err := service.GetThePagePost(userId, pageInfo)

	if err != nil {
		response.FailWithMessage("数据拉取失败", c)
		return
	}

	response.OkWithData(posts, c)
}
