package user_follow

import (
	"GO1/global"
	"GO1/models"
	"time"
)

func FollowUser(followerID int64, followeeID int64) bool {

	result := CheckFollows(followerID, followeeID)
	if !result {
		follow := models.UserFollow{
			FollowerID: followerID,
			FolloweeID: followeeID,
			CreatedAt:  time.Now(),
		}
		if err := global.DB.Create(&follow).Error; err != nil {
			return false
		}
		return true
	}
	return false
}
