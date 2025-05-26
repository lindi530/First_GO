package user

import (
	"GO1/database/redis"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func RefreshOnlineState(c *gin.Context, accessToken string) {
	claims, err := jwt.ParseToken(accessToken)

	if err != nil {
		return
	}
	redis.RefreshOnlineState(c, claims.UserId)
}
