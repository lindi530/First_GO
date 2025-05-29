package user

import (
	"GO1/global"
	"GO1/models/user"
)

func FindUser(i interface{}) user.User {
	var result user.User
	switch t := i.(type) {
	case UserNameParam:
		result = FindUserByUserName(string(t))
	case UserIdParam:
		result = FindUserByUserId(int64(t))
	default:
		break
	}

	return result
}

func FindUserByUserId(id int64) user.User {
	user := user.User{}
	global.DB.Where("user_id = ?", id).First(&user)
	return user
}

func FindUserByUserName(username string) user.User {
	user := user.User{}
	global.DB.Where("user_name = ?", username).First(&user)
	return user
}
