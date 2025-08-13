package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
	"math"
)

func GetProblemList() ([]problem_model.ProblemList, error) {

	var problems []problem_model.ProblemList

	// 查询所有题目信息
	err := global.DB.Table("problems").
		Select("id, title, level, submit_count, ac_count").
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
		rate := float64(problems[i].AcCount) / float64(problems[i].SubmitCount)
		problems[i].PassRate = math.Round(rate*10000) / 100
	}

	return problems, nil
}
