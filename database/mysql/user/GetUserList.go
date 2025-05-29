package user

import (
	"GO1/global"
	"GO1/models/user"
)

func GetUserList() []user.User {
	users := []user.User{}
	global.DB.Find(&users)
	return users
}
