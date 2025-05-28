package user_follow

import (
	"GO1/global"
	"GO1/models"
)

func UnFollowUser(followerID int64, followeeID int64) bool {
	err := global.DB.
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		Delete(&models.UserFollow{}).Error
	if err != nil {
		return false
	}
	return UpdateUnFollowUserCount(followerID, followeeID)
}
