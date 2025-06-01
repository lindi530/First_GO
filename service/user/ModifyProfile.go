package user

import (
	mysql_user "GO1/database/mysql/user"
	"GO1/global"
	"GO1/models"
	"GO1/models/user"
	"github.com/gin-gonic/gin"
)

func ModifyProfile(c *gin.Context, userId int64, profile user.UserProfile) models.HandleFuncResp {
	if result := AuthUser(c, userId); !result {
		return models.HandleFuncResp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	user := mysql_user.FindUser(mysql_user.UserIdParam(userId))
	if user.UserID == 0 {
		return models.HandleFuncResp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	global.Logger.Info("mysql ModifyProfile")
	result := mysql_user.ModifyProfile(user.UserID, profile)

	return result
}
