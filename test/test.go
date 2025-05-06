package test

import (
	"GO1/global"
	"GO1/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(router *gin.RouterGroup) {
	var users []models.User
	global.DB.Find(&users)
	fmt.Println(users)
	global.Logger.Info("日志文件")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
}
