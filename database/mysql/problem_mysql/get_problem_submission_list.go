package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func GetProblemSubmissionList(userId int64, problemId uint, submissions *[]problem_model.ProblemSubmission) error {

	err := global.DB.Table("problem_submissions AS s").
		Select("s.id, s.problem_id, s.language, s.created_at, s.score, s.exec_time, s.memory_usage, s.state").
		Where("s.user_id = ? AND s.problem_id = ?", userId, problemId).
		Order("s.created_at DESC").
		Find(submissions).Error

	//err := global.DB.
	//	Where("user_id = ? AND problem_id = ?", userId, problemId).
	//	Order("created_at DESC").
	//	//Offset((page - 1) * pageSize).                  // 分页
	//	//Limit(pageSize).
	//	Find(&submissions).Error

	return err
}
