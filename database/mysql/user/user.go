package user

import (
	"GO1/global"
	"GO1/models"
	"GO1/pkg/snowflake"
	"GO1/service/hash"
	"fmt"
	"time"
)

func Register(register models.ParamRegister) {
	userId := snowflake.Snowflake{}.GenID()
	hashPassword, err := hash.HashPassword(register.Password)
	if err != nil {
		global.Logger.Error(fmt.Sprintf("%s hash password error: %v", register.Name, err))
	}
	time := time.Now()
	user := models.User{
		UserID:     userId,
		UserName:   register.Name,
		Password:   hashPassword,
		Email:      register.Email,
		Gender:     register.Gender,
		CreateTime: time,
		UpdateTime: time,
	}
	global.DB.Create(&user)

	global.Logger.Info(user)
}
