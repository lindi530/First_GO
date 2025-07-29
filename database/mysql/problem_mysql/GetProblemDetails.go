package problem_mysql

import (
	"GO1/global"
	model "GO1/models/problem_model"
)

func GetProblemDetails(id int64, problemDetail *model.Problem) error {

	err := global.DB.First(problemDetail, "ID=?", id).Error

	return err
}
