package match_service

import (
	"GO1/database/mysql/saber_stats_mysql"
	"GO1/models/match_model"
	"GO1/models/saber_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
)

func responseMatch(user1, user2 *match_model.MatchUser, roomID string, problemID uint) {
	info1 := saber_model.SaberStat{}
	err := saber_stats_mysql.GetSaberStat(&info1, user1.UserID)
	if err != nil {
		return
	}
	info2 := saber_model.SaberStat{}
	err = saber_stats_mysql.GetSaberStat(&info2, user2.UserID)
	if err != nil {
		return
	}

	resp1 := &ws_model.MatchResponse{
		Type:      ws_model.MessageTypeMatchSuccess,
		To:        user1.UserID,
		RoomID:    roomID,
		ProblemID: problemID,
		Opponent:  user2,
		Info: ws_model.MatchInfo{
			Name:   user2.UserName,
			Avatar: user2.Avatar,

			Rating:       info2.Rating,
			Level:        info2.Level,
			Wins:         info2.Wins,
			TotalMatches: info2.TotalMatches,
		},
	}
	resp2 := &ws_model.MatchResponse{
		Type:      ws_model.MessageTypeMatchSuccess,
		To:        user2.UserID,
		RoomID:    roomID,
		ProblemID: problemID,
		Opponent:  user1,
		Info: ws_model.MatchInfo{
			Name:   user1.UserName,
			Avatar: user1.Avatar,

			Rating:       info1.Rating,
			Level:        info1.Level,
			Wins:         info1.Wins,
			TotalMatches: info1.TotalMatches,
		},
	}
	ws_service.WsHub.SendMatchData(resp1)
	ws_service.WsHub.SendMatchData(resp2)
}
