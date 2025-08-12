package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
)

func GetUserList() []user_model.User {
	users := []user_model.User{}
	global.DB.Find(&users)
	return users
}
