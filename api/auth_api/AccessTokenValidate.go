package auth_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AuthAPI) AccessTokenValidate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if _, err := jwt.ParseToken(accessToken); err != nil {
		response.FailWithCode(response.InvalidAccessToken, c)
	}

	response.OkWithMessage("成功", c)
}
