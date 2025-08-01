package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func GetProblemExamples(problemId int64, examples *[]problem_model.Example) error {

	if err := global.DB.
		Model(&problem_model.ProblemExample{}).
		Select("input", "output").
		Where("problem_id = ?", problemId).
		Scan(examples).Error; err != nil {
		return err
	}

	return nil
}
