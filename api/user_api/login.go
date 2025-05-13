package user_api

import (
	"GO1/middlewares/response"
	"GO1/models"
	"GO1/service/hash"
	userService "GO1/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (UserAPI) LoginUser(c *gin.Context) {
	// 信息校验
	user := models.ParamLogin{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithCode(http.StatusBadRequest, c)
		return
	}
	if result := hash.CheckPassword(user); result {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 登录
	userService.Login(user)

	// 返回结果
	response.OkWithMessage("登录成功！", c)
}
