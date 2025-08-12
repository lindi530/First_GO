package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func GetProblemList() ([]problem_model.ProblemListResponse, error) {

	var problems []problem_model.ProblemListResponse

	// 查询所有题目信息
	err := global.DB.Table("problems").
		Select("id, title, level").
		Find(&problems).Error

	if err != nil {
		return nil, err
	}

	if len(problems) == 0 {
		return problems, nil
	}

	//遍历题目，查询对应标签
	for i := range problems {
		var tags []string
		err = global.DB.Table("problem_tags").
			Where("problem_id = ?", problems[i].ID).
			Pluck("tag", &tags).Error
		if err != nil {
			continue
		}
		problems[i].Tags = tags
	}

	return problems, nil
}
