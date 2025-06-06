package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	auth.POST("/validate/refresh_token", api.ApiGroups.AuthAPI.RefreshTokenValidate)
	auth.GET("/validate/access_token", api.ApiGroups.AuthAPI.AccessTokenValidate)

	auth.Use(jwt.JWTAuthMiddleware())
	{

	}

}
