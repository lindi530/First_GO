package user_follow

import (
	"GO1/global"
	modelsUser "GO1/models/user"
)

func GetFollowUsers(userId int64) ([]modelsUser.User, error) {
	followUsersId, err1 := GetFollowUsersId(userId)
	userInfo, err2 := GetFollowUsersInfo(followUsersId)
	if err1 != nil {
		return []modelsUser.User{}, err1
	}
	if err2 != nil {
		return []modelsUser.User{}, err2
	}
	return userInfo, nil
}

func GetFollowUsersId(userId int64) ([]int64, error) {
	followUsersId := make([]int64, 0)
	err := global.DB.
		Table("user_follows AS a").
		Joins("JOIN user_follows AS b ON a.followee_id = b.follower_id AND b.followee_id = ? ", userId).
		Where("a.follower_id = ?", userId).
		Pluck("a.followee_id", &followUsersId).
		Error

	return followUsersId, err
}

func GetFollowUsersInfo(userId []int64) ([]modelsUser.User, error) {
	if len(userId) == 0 {
		return []modelsUser.User{}, nil
	}

	followUserList := []modelsUser.User{}
	err := global.DB.
		Where("user_id in (?)", userId).
		Find(&followUserList).
		Error
	return followUserList, err
}
