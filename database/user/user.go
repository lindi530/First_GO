package user

import (
	"GO1/global"
	"GO1/models"
	"GO1/pkg/snowflake"
	"time"
)

func CheckUserByName(name string) bool {
	user := models.User{}
	global.DB.Where("username = ?", name).First(&user)
	if user.UserID != 0 {
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

func Register(register models.Register) {
	userId := snowflake.Snowflake{}.GenID()
	time := time.Now()
	user := models.User{
		UserID:     userId,
		Username:   register.Name,
		Password:   register.Password,
		Email:      register.Email,
		Gender:     register.Gender,
		CreateTime: time,
		UpdateTime: time,
	}
	global.DB.Create(&user)

	global.Logger.Info(user)
}
