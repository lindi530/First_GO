package account

import (
	mysql "GO1/database/mysql/user"
	"GO1/database/redis"
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models"
	"GO1/service/hash"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
)

func (UserAccountAPI) Login(c *gin.Context) {
	var input struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	// 信息校验
	if err := c.ShouldBindJSON(&input); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	user := mysql.FindUser(mysql.UserNameParam(input.UserName))

	if user.UserID == 0 {
		global.Logger.Error("无效的用户信息")
		response.FailWithCode(response.InvalidLoginInfo, c)
		return
	}

	if result := hash.CheckPassword(user.Password, input.Password); !result {
		response.FailWithCode(response.InvalidLoginInfo, c)
		return
	}
	// 登录
	accessToken, refreshToken, jti := service.Login(user.UserID, user.UserName)

	redis.SaveJWTId(c, user.UserID, jti)
	redis.SaveOnlineState(c, user.UserID)

	responseUser := models.BuildUserResponse(c, user)
	// 返回结果
	response.Ok(gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         responseUser,
	}, "登录成功！", c)
}
