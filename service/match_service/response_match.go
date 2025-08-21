package match_service

import (
	"GO1/models/match_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
)

func ResponseMatch(user1, user2 *match_model.MatchUser, roomID string, problemID uint) {

	resp1 := &ws_model.MatchResponse{
		Type:      ws_model.MessageTypeMatchSuccess,
		To:        user1.UserID,
		RoomID:    roomID,
		ProblemID: problemID,
		Opponent:  user2,
	}
	resp2 := &ws_model.MatchResponse{
		Type:      ws_model.MessageTypeMatchSuccess,
		To:        user2.UserID,
		RoomID:    roomID,
		ProblemID: problemID,
		Opponent:  user1,
	}
	ws_service.WsHub.SendMatchData(resp1)
	ws_service.WsHub.SendMatchData(resp2)
}
