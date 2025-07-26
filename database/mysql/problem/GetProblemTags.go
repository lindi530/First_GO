package problem

import (
	"GO1/global"
	model "GO1/models/problem"
)

func GetProblemTags(id int64, problemDetail *model.Problem) error {
	var problemTags []string

	if err := global.DB.Model(&model.ProblemTag{}).
		Where("problem_id=?", id).
		Pluck("tag", &problemTags).Error; err != nil {
		return err
	}

	problemDetail.Tags = problemTags
	return nil
}
