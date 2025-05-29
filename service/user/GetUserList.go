package user

import (
	mysql "GO1/database/mysql/user"
	models_user "GO1/models/user"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) []models_user.UserResponse {
	users := mysql.GetUserList()
	userList := make([]models_user.UserResponse, 0, len(users))
	for _, user := range users {
		userList = append(userList, models_user.BuildUserResponse(c, user))
	}
	return userList
}
