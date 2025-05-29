package user_follow

import (
	"GO1/global"
	"GO1/models/user_follow"
)

func CheckFollows(followerID int64, followeeID int64) bool {
	follow := user_follow.UserFollow{}
	err := global.DB.Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		First(&follow).Error
	if err != nil {
		return false
	}
	return true
}
