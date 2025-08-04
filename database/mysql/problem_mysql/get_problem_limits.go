package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
	"errors"
)

func GetProblemConstraints(problemId64 int64, language string, constraint *problem_model.ProblemConstraint) error {
	problemId := uint(problemId64)
	if err := global.DB.Where("id = ?	AND language = ?", problemId, language).First(constraint).Error; err != nil {
		return errors.New("查询题目限制失败")
	}
	return nil
}
