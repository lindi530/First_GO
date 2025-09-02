package user_follow

import (
	"GO1/global"
	"GO1/models/user_model"
	"gorm.io/gorm"
)

func UpdateFollowUserCount(followerID, followeeID int64) (result bool) {

	if err := global.DB.Model(&user_model.User{}).
		Where("user_id = ?", followerID).
		UpdateColumn("following_count", gorm.Expr("following_count + ?", 1)).
		Error; err != nil {
		global.Logger.Error("更新 following_count 失败:", err)
		return false
	}
	global.Logger.Info("UpdateFollowUserCount: following_count 成功")

	if err := global.DB.Model(&user_model.User{}).
		Where("user_id = ?", followeeID).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).
		Error; err != nil {
		global.Logger.Error("更新 follower_count 失败:", err)
		return false
	}

	global.Logger.Info("UpdateFollowUserCount: follower_count 成功")
	return true
}

func UpdateUnFollowUserCount(followerID, followeeID int64) (result bool) {
	if err := global.DB.Model(&user_model.User{}).
		Where("user_id = ?", followerID).
		UpdateColumn("following_count",
			gorm.Expr("following_count - 1")).
		Error; err != nil {
		return false
	}
	if err := global.DB.Model(&user_model.User{}).
		Where("user_id = ?", followeeID).
		UpdateColumn("follower_count",
			gorm.Expr("follower_count - 1")).
		Error; err != nil {
		return false
	}
	return true
}
