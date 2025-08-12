package match_mysql

import (
	"GO1/global"
	"GO1/models/match_model"
	"GO1/models/user_model"
)

func GetMatchUser(user *match_model.MatchUser) {
	if err := global.DB.
		Model(&user_model.User{}).
		Select("user_name", "avatar", "rank").
		Where("user_id = ?", user.UserID).
		Scan(user).Error; err != nil {
		global.Logger.Error("非法用户ID")
		return
	}
	return
		
}
