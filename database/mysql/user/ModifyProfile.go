package user

import (
	"GO1/global"
	"GO1/models"
	"github.com/gin-gonic/gin"
)

func ModifyProfile(user models.User, profile models.UserProfile) gin.H {
	updates := map[string]interface{}{
		"username": profile.UserName,
		"email":    profile.Email,
	}
	err := global.DB.Model(&user).Updates(updates).Error
	if err != nil {
		return gin.H{
			"msg": "数据库保存失败",
		}
	}
	return gin.H{
		"msg": "修改成功",
	}
}
