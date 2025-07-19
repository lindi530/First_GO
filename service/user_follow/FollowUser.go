package user_follow

import "GO1/database/mysql/user_follow"

func FollowUser(followerID int64, followeeID int64) bool {
	result := user_follow.FollowUser(followerID, followeeID)
	return result
}
