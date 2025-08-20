package redis

import (
	"GO1/global"
	"GO1/models/match_model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func GetSaberRoom(roomID string) (*match_model.Room, error) {
	ctx := context.Background()
	val, err := global.RedisClient.Get(ctx, roomID).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("房间 %s 不存在", roomID)
		}
		return nil, err
	}

	var room match_model.Room
	if err := json.Unmarshal([]byte(val), &room); err != nil {
		return nil, err
	}

	return &room, nil
}
