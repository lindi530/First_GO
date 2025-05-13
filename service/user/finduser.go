package user

import (
	"GO1/global"
	"GO1/models"
)

func FindUserById(id int) models.User {
	user := models.User{}
	global.DB.Where("id = ?", id).First(&user)
	return user
}

func FindUserByUsername(username string) models.User {
	user := models.User{}
	global.DB.Where("username = ?", username).First(&user)
	return user
}
