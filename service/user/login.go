package user

import (
	"GO1/auth"
	"github.com/gin-gonic/gin"
)

func Login(username string, c *gin.Context) string {
	user := FindUserByUsername(username)
	token, _ := auth.GenerateToken(user.UserID)
	return token
}
