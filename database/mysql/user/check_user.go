package user

import (
	"GO1/global"
	"GO1/models"
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
	user := models.User{}
	global.DB.Where("username = ?", username).First(&user)
	if user.UserID != 0 {
		return true
	}
	return false
}

func CheckUserByUserId(userid int64) bool {
	user := models.User{}
	global.DB.Where("user_id = ?", userid).First(&user)
	if user.UserID != 0 {
		return true
	}
	return false
}
