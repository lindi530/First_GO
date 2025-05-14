package redis

import (
	"GO1/global"
	"github.com/go-redis/redis/v8"
)

func newClient() *redis.Client {
	addr := global.Config.RedisClient.Address
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func InitRedisClient() {
	global.RedisClient = newClient()
}
