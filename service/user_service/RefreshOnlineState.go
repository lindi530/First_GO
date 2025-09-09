package user_service

import (
	"GO1/database/redis"
	"GO1/global"
	"GO1/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
)

func RefreshOnlineState(c *gin.Context, accessToken string) {
	claims, err := jwt.ParseToken(accessToken)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			global.Logger.Warn("Token Expired")
		} else {
			global.Logger.Error("Parse Token Failed: ")
		}
		return
	}
	redis.RefreshOnlineState(c, claims.UserId)
}
