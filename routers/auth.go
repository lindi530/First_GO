package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.GET("/validate/access_token/:user_id", api.ApiGroups.AuthAPI.AccessTokenValidate)
	auth.GET("/validate/refresh_token/:user_id", api.ApiGroups.AuthAPI.RefreshTokenValidate)
}
