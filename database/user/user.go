package user

import (
	"GO1/global"
	"GO1/models"
)

func CheckUserByName(name string) bool {
	user := models.User{}
	global.DB.Where("username = ?", name).First(&user)
	if user.Id != 0 {
		return true
	}
	return false
}

func CheckUser(i interface{}) bool {
	var result bool
	switch t := i.(type) {
	case Name:
		result = CheckUserByName(string(t))
	default:

		result = false
		break
	}
	return result
}

func Register(user models.Register) {
	//global.DB.Create(&user)

}
