package problem_mysql

import (
	"GO1/global"
	model "GO1/models/problem_model"
)

func GetProblemTags(id int64, problemTags *[]string) error {

	if err := global.DB.Model(&model.ProblemTag{}).
		Where("problem_id=?", id).
		Pluck("tag", problemTags).Error; err != nil {
		return err
	}

	return nil
}
