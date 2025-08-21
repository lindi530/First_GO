package saber_service

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
)

func ExpireRoom(roomId string) (resp response.Response) {
	err := redis.ExpireRoom(roomId)

	if err != nil {
		resp.Code = 1
		resp.Message = "删除房间失败！"
		return
	}

	resp.Code = 0
	resp.Message = "删除房间成功！"
	return
}
