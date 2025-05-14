package redis

import (
	"GO1/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CheckRefreshToken(c *gin.Context, userId int64, jti string) bool {
	key := global.Config.RedisClient.KeyPrefix + jti
	val, err := global.RedisClient.Get(c, key).Result()
	if err != nil || val != strconv.FormatInt(userId, 10) {
		return false
	}
	return true
}
