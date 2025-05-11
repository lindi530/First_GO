package user_api

import (
	"GO1/global"
	"GO1/models"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UsersAPI struct{}

func (UsersAPI) UsersInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "users_info",
	})
}

func (UsersAPI) RegisterUser(c *gin.Context) {
	// 数据校验
	user := models.Register{}
	if err := c.ShouldBindJSON(&user); err != nil {
		global.Logger.Info("注册数据错误！", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"data":    user,
		})
		return
	}

	// 调用服务
	result := service.Register(user)

	if !result {
		c.JSON(http.StatusOK, "已存在该用户！")
		return
	}

	// 响应结果
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功！",
	})
}
