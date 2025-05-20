package user_api

import (
	mysql "GO1/database/mysql/user"
	"GO1/middlewares/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	if result := mysql.CheckUser(mysql.UserIdParam(userId)); result == false {
		response.FailWithMessage("该用户不存在！", c)
		return
	}
	if result := mysql.DeleteUser(userId); result == false {
		response.FailWithMessage("删除失败！", c)
		return
	}
	response.OkWithData(gin.H{
		"userId": userId,
		"msg":    "success delete",
	}, c)
}
