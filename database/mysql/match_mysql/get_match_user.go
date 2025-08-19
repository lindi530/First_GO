package match_mysql

import (
	"GO1/global"
	"GO1/models/match_model"
	"GO1/models/saber_model"
	"GO1/models/user_model"
	"time"
)

func GetMatchUser(user *match_model.MatchUser, userid int64) {
	var u user_model.User
	if err := global.DB.
		Table("users").
		Select("user_name", "avatar").
		Where("user_id = ?", userid).
		Scan(&u).Error; err != nil {
		global.Logger.Error("非法用户ID")
		return
	}
	global.Logger.Info("GetMatchUser: ", user)

	var stat saber_model.SaberStat
	if err := global.DB.
		Where("user_id = ?", userid).
		Attrs(saber_model.SaberStat{ // 你想要的初始值，而不是零
			Rating:       1500,
			Level:        "黄铜",
			Wins:         0,
			TotalMatches: 0,
			CreatedAT:    time.Now(),
			UpdatedAT:    time.Now(),
		}).
		FirstOrCreate(&stat).
		Error; err != nil {
		global.Logger.Error("非法用户")
		return
	}

	user.UserID = userid
	user.UserName = u.UserName
	user.Avatar = user_model.GetAvatarPath(u.Avatar)
	user.Rating = stat.Rating
	user.Level = stat.Level
	user.Wins = stat.Wins
	user.TotalMatches = stat.TotalMatches

	global.Logger.Info("GetMatchUser: ", user)
	return
}
