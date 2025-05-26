package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	service_user "GO1/service/user"
	"github.com/gin-gonic/gin"
	"strings"
)

func (UserAPI) RefreshOnlineState(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	global.Logger.Info("accessToken: ", accessToken)

	service_user.RefreshOnlineState(c, accessToken)
	global.Logger.Info("refresh online state success")
	response.OkWithMessage("", c)
}
