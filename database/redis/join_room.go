package redis

import (
	"GO1/global"
	"GO1/models/match_model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func GetJoinRoom(roomId string, userId int64) (*match_model.Room, error) {
	ctx := context.Background()

	// 1. 从 Redis 取出房间信息
	val, err := global.RedisClient.Get(ctx, roomId).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("房间不存在")
	} else if err != nil {
		return nil, err
	}

	// 2. 反序列化 JSON
	var room match_model.Room
	if err := json.Unmarshal([]byte(val), &room); err != nil {
		return nil, err
	}

	// 3. 判断房间状态
	if room.User2ID != 0 {
		return nil, fmt.Errorf("房间已满")
	}

	return &room, nil
}

func JoinRoom(room *match_model.Room, userId int64, problemId uint) error {
	// 4. 加入用户
	room.User2ID = userId
	room.ProblemID = problemId
	room.Status = "ready"

	// 5. 重新存入 Redis，保持 TTL 不变（30 分钟）
	jsonData, err := json.Marshal(&room)
	if err != nil {
		return err
	}

	// ⚠️ 这里需要用 Set + TTL，否则会覆盖原来的过期时间
	ctx := context.Background()
	ttl, _ := global.RedisClient.TTL(ctx, room.RoomID).Result()
	if ttl <= 0 {
		ttl = 30 * time.Minute
	}

	err = global.RedisClient.Set(ctx, room.RoomID, jsonData, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}
