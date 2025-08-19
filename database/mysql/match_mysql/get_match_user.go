package match_mysql

import (
	"GO1/global"
	"GO1/models/match_model"
)

func GetMatchUser(user *match_model.MatchUser, userid int64) {
	if err := global.DB.
		Table("users").
		Select("user_name", "avatar").
		Where("user_id = ?", userid).
		Scan(user).Error; err != nil {
		global.Logger.Error("非法用户ID")
		return
	}
	global.Logger.Info("GetMatchUser: ", user)
	if err := global.DB.
		Table("saber_stats").
		Where("user_id = ?", userid).
		Pluck("rating", &user.Rating).
		Error; err != nil {
		global.Logger.Error("非法用户")
		return
	}

	user.UserID = userid
	global.Logger.Info("GetMatchUser: ", user)
	return
}
