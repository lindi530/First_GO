package users_api

import "github.com/gin-gonic/gin"

func (UsersAPI) UsersInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users_info",
	})
}
