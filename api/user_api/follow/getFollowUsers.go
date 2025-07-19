package follow

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/user"
	"github.com/gin-gonic/gin"
)

func (UserFollowAPI) GetFollowUsers(c *gin.Context) {
	// 数据处理
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	// 服务
	resp := user.GetFollowUsers(c, userId)

	response.Ok(resp.Data, resp.Message, c)
}
