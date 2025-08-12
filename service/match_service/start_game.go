package match_service

import (
	"GO1/database/redis"
	"GO1/global"
	"GO1/models/match_model"
)

func startGame(user1, user2 *match_model.MatchUser) {
	// TODO: 创建房间，分配题目，通知双方 WebSocket（或 HTTP 回调）
	global.Logger.Errorf("Start Game: %s vs %s", user1.Username, user2.Username)
	problemID := SelectProblemID(user1.Rank, user2.Rank)
	room, err := redis.CreateRoom(user1.UserID, user2.UserID, problemID)
	if err != nil {
		// ws 通知客户端
		global.Logger.Error("服务器创建房间失败！！！")
		return
	}
	global.Logger.Info("匹配成功！！！")

	// 发送 WebSocket 通知双方进入房间
	responseMatch(user1, user2, room.RoomID, problemID)
}
