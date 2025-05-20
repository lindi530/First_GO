package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
)

func (UserAPI) GetUserList(c *gin.Context) {
	// 数据校验

	// 服务
	userList := service.GetUserList()
	// 相应
	global.Logger.Info("UserList: ", userList)
	response.OkWithData(gin.H{
		"userList": userList,
	}, c)
}
