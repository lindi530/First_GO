package post_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/post"
	service "GO1/service/user/post"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (PostAPI) GetAllPost(c *gin.Context) {
	// 数据校验
	// 服务
	posts, _ := service.GetAllPost(c)
	fmt.Println(posts[0])
	// 响应
	response.OkWithData(posts, c)
}

func (PostAPI) GetOnePost(c *gin.Context) {
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	result := service.GetOnePost(c, postId)

	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
		return
	}
	response.OkWithData(result.Data, c)
}

func (PostAPI) GetThePagePost(c *gin.Context) {
	var pageInfo post.PageInfo
	if err := c.ShouldBind(&pageInfo); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}
	global.Logger.Infof("GetThePagePost pageInfo: %v", pageInfo)
	posts, err := service.GetThePagePost(c, pageInfo)

	if err != nil {
		response.FailWithMessage("数据拉取失败", c)
		return
	}

	response.OkWithData(posts, c)
}
