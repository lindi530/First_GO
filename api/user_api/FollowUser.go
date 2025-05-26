package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/service/user_follow"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) FollowUser(c *gin.Context) {
	followerID := GetUserIdFromToken(c.GetHeader("Authorization"))
	followeeID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	global.Logger.Info(fmt.Sprintf("Following follower %d, followee %d", followerID, followeeID))
	if followerID == followeeID {
		response.FailWithMessage("不允许自己关注自己", c)
		return
	}

	result := user_follow.FollowUser(followerID, followeeID)

	if !result {
		response.OkWithMessage("关注失败", c)
		return
	}

	response.OkWithMessage("关注成功", c)
}
