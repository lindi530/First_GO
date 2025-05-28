package follow

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	user_follow "GO1/service/user_follow"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserFollowAPI) UnFollowUser(c *gin.Context) {
	followerID := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	followeeID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if followerID == followeeID {
		response.FailWithMessage("不存在自己关注自己", c)
		return
	}

	result := user_follow.UnFollowUser(followerID, followeeID)

	if !result {
		response.OkWithMessage("关注失败", c)
		return
	}

	response.OkWithMessage("关注成功", c)
}
