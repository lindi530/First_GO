package user_follow

import (
	"GO1/global"
	"GO1/models/user_follow"
)

func UnFollowUser(followerID int64, followeeID int64) bool {
	err := global.DB.
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		Delete(&user_follow.UserFollow{}).Error
	if err != nil {
		return false
	}
	return UpdateUnFollowUserCount(followerID, followeeID)
}
