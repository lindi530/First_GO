package saber_stats_mysql

import (
	"GO1/global"
	"GO1/models/saber_model"
	"time"
)

func GetSaberStat(stat *saber_model.SaberStat, userid int64) error {
	// 手写
	//err := global.DB.Where("user_id = ?", userid).First(stat).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	*stat = saber_model.SaberStat{UserID: userid} // 用解引用修改原值
	//	err = global.DB.Create(stat).Error
	//}
	//return err

	// gorm 写法
	err := global.DB.
		Where(saber_model.SaberStat{UserID: userid}).
		Attrs(saber_model.SaberStat{ // 你想要的初始值，而不是零
			Rating:       1500,
			Level:        "黄铜",
			Wins:         0,
			TotalMatches: 0,
			CreatedAT:    time.Now(),
			UpdatedAT:    time.Now(),
		}).
		FirstOrCreate(stat).
		Error
	return err
}
