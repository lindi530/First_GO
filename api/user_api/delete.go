package user_api

import (
	"GO1/middlewares/response"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) DeleteUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	user := service.FindUserById(userid)
	if user.Username == "" {
		response.FailWithMessage("该用户不存在！", c)
		return
	}
	if result := service.DeleteUser(userid); result != true {
		response.FailWithMessage("删除失败！", c)
		return
	}
	response.OkWithData(gin.H{
		"user": user,
		"msg":  "success delete",
	}, c)
}
