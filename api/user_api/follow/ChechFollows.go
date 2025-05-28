package follow

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	service_user "GO1/service/user_follow"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserFollowAPI) CheckFollows(c *gin.Context) {
	followerID := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
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
