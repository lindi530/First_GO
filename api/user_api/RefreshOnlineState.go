package user_api

import (
	"GO1/middlewares/response"
	service_user "GO1/service/user_service"
	"github.com/gin-gonic/gin"
	"strings"
)

func (UserAPI) RefreshOnlineState(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	service_user.RefreshOnlineState(c, accessToken)

	response.OkWithMessage("", c)
}
