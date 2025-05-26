package user_follow

import "GO1/database/mysql/user_follow"

func UnFollowUser(followerID int64, followeeID int64) bool {
	return user_follow.UnFollowUser(followerID, followeeID)
}
