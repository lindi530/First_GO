package saber_service

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
)

func InviteFriend(userid int64) (resp response.Response) {
	room, err := redis.CreateInviteRoom(userid)
	if err != nil {
		resp.Code = 1
		resp.Message = "创建房间失败！"
		return
	}

	resp.Code = 0
	resp.Data = room
	resp.Message = "创建房间成功！"

	return resp
}
