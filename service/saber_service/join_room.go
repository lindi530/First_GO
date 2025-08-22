package saber_service

import (
	"GO1/database/mysql/match_mysql"
	"GO1/database/redis"
	"GO1/middlewares/response"
	"GO1/models/match_model"
	"GO1/service/match_service"
)

func JoinRoom(roomId string, userId int64) (resp response.Response) {
	room, err := redis.GetJoinRoom(roomId, userId)

	if err != nil {
		resp.Code = 1
		resp.Message = "房间获取失败！"
		return
	}

	user1 := match_model.MatchUser{}
	match_mysql.GetMatchUser(&user1, room.User1ID)
	user2 := match_model.MatchUser{}
	match_mysql.GetMatchUser(&user2, userId)

	problemId := match_service.SelectProblemID(user1.Rating, user2.Rating)

	err = redis.JoinRoom(room, userId, problemId)
	if err != nil {
		resp.Code = 1
		resp.Message = "加入房间失败！"
		return
	}

	match_service.ResponseMatch(&user1, &user2, roomId, problemId)

	resp.Code = 0
	resp.Message = "加入房间成功！"
	return
}
