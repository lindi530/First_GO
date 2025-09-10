package redis

import (
	"GO1/global"
	"github.com/go-redis/redis/v8"
	"os"
)

func newClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = global.Config.RedisClient.Address
	}
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func InitRedisClient() {
	global.RedisClient = newClient()
}
