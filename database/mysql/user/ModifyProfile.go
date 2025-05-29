package user

import (
	"GO1/global"
	"GO1/models"
	"GO1/models/user"
)

func ModifyProfile(user user.User, profile user.UserProfile) (rsp models.HandleFuncResp) {
	updates := map[string]interface{}{
		"user_name": profile.UserName,
		"email":     profile.Email,
		"quote":     profile.Quote,
	}
	global.Logger.Info(profile)
	global.Logger.Info(user)
	err := global.DB.Model(&user).Updates(updates).Error
	if err != nil {
		return models.HandleFuncResp{
			Msg: "数据库存储失败",
			Ok:  false,
		}
	}
	return models.HandleFuncResp{
		Msg: "修改成功",
		Ok:  true,
	}
}
