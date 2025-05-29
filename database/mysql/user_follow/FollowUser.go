package user_follow

import (
	"GO1/global"
	"GO1/models/user_follow"
	"time"
)

func FollowUser(followerID int64, followeeID int64) bool {

	result := CheckFollows(followerID, followeeID)
	if !result {
		follow := user_follow.UserFollow{
			FollowerID: followerID,
			FolloweeID: followeeID,
			CreatedAt:  time.Now(),
		}
		if err := global.DB.Create(&follow).Error; err != nil {
			return false
		}
		return UpdateFollowUserCount(followerID, followeeID)
	}
	return false
}
