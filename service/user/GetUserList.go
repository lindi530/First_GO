package user

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
)

func GetUserList() []models.UserResponse {
	users := mysql.GetUserList()
	userList := make([]models.UserResponse, 0, len(users))
	for _, user := range users {
		userList = append(userList, models.BuildUserResponse(user))
	}
	return userList
}
