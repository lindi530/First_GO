package user

import (
	"GO1/database/mysql/user_mysql"
	"GO1/models/user_model"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) []user_model.UserResponse {
	users := user_mysql.GetUserList()
	userList := make([]user_model.UserResponse, 0, len(users))
	for _, user := range users {
		userList = append(userList, user_model.BuildUserResponse(c, user))
	}
	return userList
}
