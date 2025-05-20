package user

import (
	"GO1/global"
	"GO1/models"
)

func GetUserList() []models.User {
	users := []models.User{}
	global.DB.Find(&users)
	return users
}
