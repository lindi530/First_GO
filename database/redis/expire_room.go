package redis

import (
	"GO1/global"
	"context"
)

func ExpireRoom(roomId string) error {

	ctx := context.Background()

	err := global.RedisClient.Del(ctx, roomId).Err()
	if err != nil {
		global.Logger.Error("删除房间失败:", err)
		return err
	}

	return nil
}
