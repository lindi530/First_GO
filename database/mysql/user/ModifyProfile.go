package user

import (
	"GO1/global"
	"GO1/models"
)

func ModifyProfile(user models.User, profile models.UserProfile) (rsp models.HandleFuncRsp) {
	updates := map[string]interface{}{
		"user_name": profile.UserName,
		"email":     profile.Email,
		"quote":     profile.Quote,
	}
	global.Logger.Info(profile)
	global.Logger.Info(user)
	err := global.DB.Model(&user).Updates(updates).Error
	if err != nil {
		return models.HandleFuncRsp{
			Msg: "数据库存储失败",
			Ok:  false,
		}
	}
	return models.HandleFuncRsp{
		Msg: "修改成功",
		Ok:  true,
	}
}
