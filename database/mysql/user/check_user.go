package user

import (
	"GO1/global"
	"GO1/models/user"
)

func CheckUser(i interface{}) bool {
	var result bool
	switch t := i.(type) {
	case UserNameParam:
		result = CheckUserByUserName(string(t))
	case UserIdParam:
		result = CheckUserByUserId(int64(t))
	default:
		result = false
		break
	}
	return result
}

func CheckUserByUserName(username string) bool {
	user := user.User{}

	err := global.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		global.Logger.Error(err)
		return false
	}
	//if user.UserID {
	//	return true
	//}
	return user.UserID != 0
}

func CheckUserByUserId(userid int64) bool {
	user := user.User{}
	global.DB.Where("user_id = ?", userid).First(&user)
	if user.UserID != 0 {
		return true
	}
	return false
}
