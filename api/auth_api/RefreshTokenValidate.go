package auth_api

import (
	"GO1/middlewares/response"
	service_auth "GO1/service/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (AuthAPI) RefreshTokenValidate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if authHeader == "" {
		response.FailWithCode(response.InvalidRefreshToken, c)
		return
	}

	result := service_auth.RefreshTokenValidate(userId, authHeader)

	if result == false {
		response.FailWithCode(response.InvalidRefreshToken, c)
		return
	}

	response.OkWithMessage("成功", c)
}
