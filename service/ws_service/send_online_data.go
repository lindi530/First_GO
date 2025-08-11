package ws_service

import (
	"GO1/database/mysql/user_follow"
	"GO1/database/redis"
	"GO1/global"
	"GO1/models/ws_model"
	"context"
	"encoding/json"
	"fmt"
)

func (h *Hub) SendOnlineData(userId int64, onlineState bool) {
	var users []int64
	getSendUsers(&users, userId)

	onlineStatus := &ws_model.OnlineStatus{
		Type:        "online_status",
		From:        userId,
		OnlineState: onlineState,
	}

	onlineStateBytes, _ := json.Marshal(onlineStatus)

	global.Logger.Info(len(users))
	for _, user := range users {
		if receiver, ok := h.clients[user]; ok {
			receiver.Send <- onlineStateBytes

			global.Logger.Info(fmt.Sprintf("Client %d online state updated with hub", user))
		}
	}
}

func getSendUsers(users *[]int64, userId int64) {
	followUsersId, err1 := user_follow.GetFollowUsersId(userId)
	if err1 != nil {
		global.Logger.Error(err1)
		return
	}

	ctx := context.Background()

	for _, uid := range followUsersId {
		if !redis.GetOnlineState(ctx, uid) {
			continue
		}
		*users = append(*users, uid)
	}
}
