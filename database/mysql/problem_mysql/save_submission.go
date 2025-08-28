package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func SaveSubmission(problemSubmission *problem_model.ProblemSubmission) {
	if err := global.DB.Create(&problemSubmission).Error; err != nil {
		global.Logger.Error(err)
		return
	}
}
