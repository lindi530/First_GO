package user_follow

import (
	mysql_user "GO1/database/mysql/user_follow"
)

func CheckFollows(followerID int64, followeeID int64) bool {
	return mysql_user.CheckFollows(followerID, followeeID)
}
