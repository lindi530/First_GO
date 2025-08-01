package user

import (
	"GO1/database/mysql/user_follow"
	"GO1/middlewares/response"
	models_user "GO1/models/user"
	"github.com/gin-gonic/gin"
)

func GetFollowUsers(c *gin.Context, userId int64) response.Response {
	resp := response.Response{}

	followUserList, err := user_follow.GetFollowUsers(userId)

	if err != nil {
		resp.Code = 1
		resp.Message = "关注列表获取错误"
	} else {
		resp.Code = 0
		resp.Message = "关注列表获取成功"
	}

	followUserListResp := make([]models_user.UserResponse, 0, len(followUserList))
	for _, user := range followUserList {
		followUserListResp = append(followUserListResp, models_user.BuildUserResponse(c, user))
	}

	resp.Data = gin.H{
		"followUserList": followUserListResp,
	}

	return resp
}
