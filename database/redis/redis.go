package redis

import (
	"GO1/global"
	"github.com/gin-gonic/gin"
	"time"
)

func SaveJWTId(c *gin.Context, userId int64, jti string) {
	keyPrefix := global.Config.RedisClient.KeyPrefix
	key := keyPrefix + jti
	global.RedisClient.Set(c, key, userId, 7*24*time.Hour)
}

func DeleteJWTId(c *gin.Context, jti string) {
	key := global.Config.RedisClient.KeyPrefix + jti
	global.RedisClient.Del(c, key)
}
