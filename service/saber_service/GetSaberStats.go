package saber_service

import (
	"GO1/database/mysql/saber_stats_mysql"
	"GO1/middlewares/response"
	"GO1/models/saber_model"
)

func GetSaberStat(userid int64) (resp response.Response) {
	var userSaberStat saber_model.SaberStat
	err := saber_stats_mysql.GetSaberStat(&userSaberStat, userid)
	if err != nil {
		resp.Data = 1
		resp.Message = "天人数据获取失败"
		return
	}

	resp.Data = 0
	resp.Data = userSaberStat
	return
}
