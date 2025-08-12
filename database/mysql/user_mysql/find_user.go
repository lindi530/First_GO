package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
)

func FindUser(i interface{}) user_model.User {
	var result user_model.User
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

func FindUserByUserId(id int64) user_model.User {
	user := user_model.User{}
	global.DB.Where("user_id = ?", id).First(&user)
	return user
}

func FindUserByUserName(username string) user_model.User {
	user := user_model.User{}
	global.DB.Where("user_name = ?", username).First(&user)
	return user
}
