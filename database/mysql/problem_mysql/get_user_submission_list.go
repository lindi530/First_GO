package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_submission_model"
)

func GetUserSubmissionList(submissions *[]problem_submission_model.UserSubmission, userId int64, offset, limit int) error {
	err := global.DB.
		Table("problem_submissions AS s").
		Select("s.id, s.code, s.created_at, s.problem_id, p.title").
		Joins("LEFT JOIN problems p ON s.problem_id = p.id").
		Where("s.user_id = ?", userId).
		Order("id desc"). // 按 id 降序，最新的排前面
		Offset(offset). // 跳过前 offset 条
		Limit(limit). // 取 limit 条
		Find(submissions).Error
	return err
}
