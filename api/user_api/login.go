package user_api

import (
	"GO1/middlewares/response"
	"GO1/service/hash"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (UserAPI) LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 信息校验
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithCode(http.StatusBadRequest, c)
		return
	}
	if result := hash.CheckPassword(input.Username, input.Password); result {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	// 登录
	token := service.Login(input.Username, c)

	// 返回结果
	response.Ok(gin.H{
		"username": input.Username,
		"token":    token,
	}, "登录成功！", c)
}
