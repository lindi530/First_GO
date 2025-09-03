package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
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
	user := user_model.User{}

	err := global.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		global.Logger.Error(err)
		return false
	}
	//if user_mysql.UserID {
	//	return true
	//}
	return user.UserID != 0
}

func CheckUserByUserId(userid int64) bool {

	var count int64
	err := global.DB.
		Model(&user_model.User{}).
		Where("user_id = ?", userid).
		Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}
