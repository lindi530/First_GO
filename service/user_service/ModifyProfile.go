package user_service

import (
	"GO1/database/mysql/user_mysql"
	"GO1/global"
	"GO1/models"
	"GO1/models/user_model"
	"github.com/gin-gonic/gin"
)

func ModifyProfile(c *gin.Context, userId int64, profile user_model.UserProfile) models.HandleFuncResp {
	if result := AuthUser(c, userId); !result {
		return models.HandleFuncResp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	user := user_mysql.FindUser(user_mysql.UserIdParam(userId))
	if user.UserID == 0 {
		return models.HandleFuncResp{
			Msg: "无权限",
			Ok:  false,
		}
	}
	global.Logger.Info("mysql ModifyProfile")
	result := user_mysql.ModifyProfile(user.UserID, profile)

	return result
}
