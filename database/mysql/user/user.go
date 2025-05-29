package user

import (
	"GO1/global"
	"GO1/models/user"
	"GO1/pkg/snowflake"
	"GO1/service/hash"
	"fmt"
	"time"
)

func Register(register user.ParamRegister) {

	userId := snowflake.Snowflake{}.GenID()
	hashPassword, err := hash.HashPassword(register.Password)

	if err != nil {
		global.Logger.Error(fmt.Sprintf("%s hash password error: %v", register.Name, err))
	}
	time := time.Now()
	user := user.User{
		UserID:         userId,
		UserName:       register.Name,
		Password:       hashPassword,
		Email:          register.Email,
		Gender:         global.Config.Settings.User.Gender,
		Quote:          global.Config.Settings.User.Quote,
		Avatar:         global.Config.Settings.User.Avatar,
		FollowerCount:  0,
		FollowingCount: 0,
		CreateTime:     time,
		UpdateTime:     time,
	}
	global.DB.Create(&user)

	global.Logger.Info(user)
}
