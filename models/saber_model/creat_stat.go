package saber_model

import "time"

func CreateSaberStat(userid int64) SaberStat {
	saberStat := SaberStat{
		UserID:       userid,
		Rating:       1500,
		Level:        "黄铜",
		Wins:         0,
		TotalMatches: 0,
		CreatedAT:    time.Now(),
		UpdatedAT:    time.Now(),
	}
	return saberStat
}
