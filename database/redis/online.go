package redis

import (
	"GO1/global"
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var (
	onlineStatePrefix = "user:online:"
	duration          = time.Duration(15)
)

func SaveOnlineState(c *gin.Context, userId int64) {
	key := onlineStatePrefix + strconv.FormatInt(userId, 10)
	global.RedisClient.Set(c, key, 1, time.Minute*duration)
}

func RefreshOnlineState(c *gin.Context, userId int64) {
	key := onlineStatePrefix + strconv.FormatInt(userId, 10)
	ok, err := global.RedisClient.Expire(c, key, time.Minute*duration).Result()
	if err != nil {
		return
	}
	if !ok {
		SaveOnlineState(c, userId)
	}
}

func GetOnlineState(ctx interface{}, userId int64) bool {
	key := onlineStatePrefix + strconv.FormatInt(userId, 10)
	var redisCtx context.Context

	switch c := ctx.(type) {
	case *gin.Context:
		redisCtx = c.Request.Context()
	case context.Context:
		redisCtx = c
	default:
		global.Logger.Error("无效的上下文类型")
		return false
	}
	result, err := global.RedisClient.Exists(redisCtx, key).Result()
	if err != nil {
		global.Logger.Error("Redis 链接失败")
		return false
	}
	return result == 1
}

func DeleteOnlineState(c *gin.Context, userId int64) {

	key := onlineStatePrefix + strconv.FormatInt(userId, 10)

	res, err := global.RedisClient.Del(c, key).Result()
	if err != nil {
		global.Logger.Error("Redis 删除失败:", err)
	} else {
		global.Logger.Infof("Redis 删除成功，删除了 %d 个 key", res)
	}
}
