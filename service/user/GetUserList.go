package user

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) []models.UserResponse {
	users := mysql.GetUserList()
	userList := make([]models.UserResponse, 0, len(users))
	for _, user := range users {
		userList = append(userList, models.BuildUserResponse(c, user))
	}
	return userList
}
