package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
	lo "github.com/samber/lo"
)

func GetProblemList(userID int64) ([]problem_model.ProblemList, error) {

	var problems []problem_model.ProblemList

	// 查询所有题目信息
	err := global.DB.Table("problems p").
		Select(`
			p.id, 
			p.title, 
			p.level, 
			p.submit_count, 
			p.ac_count,
			CASE 
				WHEN p.submit_count = 0 THEN 0
				ELSE ROUND(p.ac_count * 1.0 / p.submit_count * 100, 2)
			END AS pass_rate,
			CASE 
				WHEN uap.problem_id IS NULL THEN false 
				ELSE true 
			END AS accepted
		`).
		Joins("LEFT JOIN user_ac_problems uap ON uap.problem_id = p.id AND uap.user_id = ?", userID).
		Scan(&problems).Error

	if err != nil {
		return nil, err
	}

	var tagRows []struct {
		ProblemID uint
		Tag       string
	}
	err = global.DB.Table("problem_tags").
		Where("problem_id IN (?)", lo.Map(problems, func(p problem_model.ProblemList, _ int) uint {
			return p.ID
		})).
		Find(&tagRows).Error

	// 用 map 聚合
	tagMap := make(map[uint][]string)
	for _, row := range tagRows {
		tagMap[row.ProblemID] = append(tagMap[row.ProblemID], row.Tag)
	}
	for i := range problems {
		problems[i].Tags = tagMap[problems[i].ID]
	}

	return problems, nil
}
