package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
	"gorm.io/gorm"
)

func UpdateAcCount(problemID uint) {
	global.DB.
		Model(&problem_model.Problem{}).
		Where("id = ?", problemID).
		UpdateColumn("ac_count", gorm.Expr("ac_count + ?", 1))
}

func UpdateSubmitCount(problemID uint) {
	global.DB.
		Model(&problem_model.Problem{}).
		Where("id = ?", problemID).
		UpdateColumn("submit_count", gorm.Expr("submit_count + ?", 1))
}
