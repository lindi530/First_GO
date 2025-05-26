package user_api

import (
	"GO1/middlewares/response"
	service_user "GO1/service/user_follow"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) CheckFollows(c *gin.Context) {
	followerID := GetUserIdFromToken(c.GetHeader("Authorization"))
	followeeID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if followerID == followeeID {
		response.FailWithMessage("不允许自己关注自己", c)
		return
	}

	result := service_user.CheckFollows(followerID, followeeID)

	response.OkWithData(gin.H{
		"isFollowing": result,
	}, c)
}
