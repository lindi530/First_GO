package user

import (
	"GO1/global"
	"GO1/models"
	"GO1/pkg/snowflake"
	"GO1/service/hash"
	"fmt"
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

func Register(register models.ParamRegister) {
	userId := snowflake.Snowflake{}.GenID()
	hashPassword, err := hash.HashPassword(register.Password)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("%s hash password error: %v", register.Name, err))
	}
	time := time.Now()
	user := models.User{
		UserID:     userId,
		Username:   register.Name,
		Password:   hashPassword,
		Email:      register.Email,
		Gender:     register.Gender,
		CreateTime: time,
		UpdateTime: time,
	}
	global.DB.Create(&user)
	
	global.Logger.Info(user)
}
