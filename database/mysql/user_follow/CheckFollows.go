package user_follow

import (
	"GO1/global"
	"GO1/models"
)

func CheckFollows(followerID int64, followeeID int64) bool {
	follow := models.UserFollow{}
	err := global.DB.Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		First(&follow).Error
	if err != nil {
		return false
	}
	return true
}
