package saber_service

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
	"GO1/service/ws_service"
)

func JoinRoom(roomId string, userId int64) (resp response.Response) {
	room, err := redis.JoinRoom(roomId, userId)

	if err != nil {
		resp.Code = 1
		resp.Message = "加入房间失败！"
		return
	}

	ws_service.WsHub.SendSaberResult()

	resp.Code = 0
	resp.Message = "加入房间成功！"
	return
}
