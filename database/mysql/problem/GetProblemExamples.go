package problem

import (
	"GO1/global"
	model "GO1/models/problem"
)

func GetProblemExamples(id int64, problemDetail *model.Problem) error {
	var examples []model.Example

	if err := global.DB.
		Model(&model.ProblemExample{}).
		Select("input", "output").
		Where("problem_id = ?", id).
		Scan(&examples).Error; err != nil {
		return err
	}
	problemDetail.Examples = examples

	return nil
}
