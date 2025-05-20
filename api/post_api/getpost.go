package post_api

import (
	"GO1/middlewares/response"
	service "GO1/service/user/post"
	"github.com/gin-gonic/gin"
)

func (PostAPI) GetAllPost(c *gin.Context) {
	// 数据校验
	// 服务
	posts, _ := service.GetAllPost()

	// 响应
	response.OkWithData(posts, c)
}
