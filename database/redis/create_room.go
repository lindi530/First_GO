package redis

import (
	"GO1/global"
	"GO1/models/match_model"
	"context"
	"encoding/json"
	"strconv"
	"time"
)

func CreateRoom(user1ID, user2ID int64, problemID uint) (*match_model.Room, error) {
	roomID := "room_" + strconv.FormatInt(time.Now().UnixNano(), 10)

	room := match_model.Room{
		RoomID:    roomID,
		User1ID:   user1ID,
		User2ID:   user2ID,
		ProblemID: problemID,
		Status:    "waiting",
	}

	jsonData, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = global.RedisClient.Set(ctx, roomID, jsonData, 30*time.Minute).Err()
	if err != nil {
		return nil, err
	}

	return &room, nil
}
