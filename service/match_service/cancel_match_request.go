package match_service

import (
	"GO1/database/mysql/match_mysql"
	"GO1/middlewares/response"
	"GO1/models/match_model"
)

func CancelMatch(userid int64) (resp response.Response) {
	user := match_model.MatchUser{}
	match_mysql.GetMatchUser(&user, userid)

	resp.Code = 0
	if user == *waitingUser {
		waitingUser = nil
		resp.Message = "取消匹配成功"
	} else {
		resp.Message = "暂无匹配信息"
	}

	return
}
