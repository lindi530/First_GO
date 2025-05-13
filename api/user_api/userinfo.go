package user_api

import "github.com/gin-gonic/gin"

func (UserAPI) UserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users_info",
	})
}
