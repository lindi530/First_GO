package user

import (
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
	"github.com/gin-gonic/gin"
)

func ModifyProfile(c *gin.Context, userId int64, profile models.UserProfile) models.HandleFuncRsp {
	if result := AuthUser(c, userId); !result {
		return models.HandleFuncRsp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	user := mysql_user.FindUser(mysql_user.UserIdParam(userId))
	if user.UserID == 0 {
		return models.HandleFuncRsp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	result := mysql_user.ModifyProfile(user, profile)

	return result
}
