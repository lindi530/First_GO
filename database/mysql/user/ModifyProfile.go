package user

import (
	"GO1/global"
	"GO1/models"
	models_user "GO1/models/user"
)

func ModifyProfile(userId int64, profile models_user.UserProfile) (resp models.HandleFuncResp) {
	updates := map[string]interface{}{
		"user_name": profile.UserName,
		"email":     profile.Email,
		"quote":     profile.Quote,
	}

	err := global.DB.Model(&models_user.User{}).
		Where("user_id = ?", userId).
		Updates(updates).Error

	if err != nil {
		global.Logger.Error(err)
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
