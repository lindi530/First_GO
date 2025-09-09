package saber_service

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
	"GO1/models/match_model"
	"GO1/models/problem_model"
	"GO1/models/saber_model"
	"GO1/models/ws_model"
	"GO1/service/problem_service"
	"GO1/service/ws_service"
)

func SaberSubmit(userid int64, submit *saber_model.SaberSubmit) (resp response.Response) {

	message := &ws_model.EditStatus{
		Type: ws_model.MessageTypeEditStatus,
		To:   userid,
	}
	ws_service.WsHub.SendEditData(message, "Pending")

	room, err := redis.GetSaberRoom(submit.RoomId)
	if err != nil {
		resp.Code = 1
		resp.Message = "不存在该游戏房间"
		return
	}

	codeSubmission := problem_model.CodeSubmission{
		Code:      submit.Code,
		Language:  submit.Language,
		ProblemID: room.ProblemID,
	}

	resp = problem_service.SubmitCode(userid, codeSubmission, message)

	if resp.Code == 0 {
		responseSaberResult(userid, room)
	}

	return
}

func responseSaberResult(userid int64, room *match_model.Room) {
	loserID := room.User1ID
	if loserID == userid {
		loserID = room.User2ID
	}
	// 通知对手我已经ac了
	saberResultLoser := ws_model.SaberResult{
		Type: ws_model.MessageTypeSaberResult,
		To:   loserID,
		Win:  false,
	}
	ws_service.WsHub.SendSaberResult(&saberResultLoser)

	winnerID := room.User1ID
	if loserID == room.User1ID {
		winnerID = room.User2ID
	}

	saberResultWinner := ws_model.SaberResult{
		Type: ws_model.MessageTypeSaberResult,
		To:   winnerID,
		Win:  true,
	}
	ws_service.WsHub.SendSaberResult(&saberResultWinner)
}
