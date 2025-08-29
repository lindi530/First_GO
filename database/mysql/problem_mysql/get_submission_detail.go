package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_model"
)

func GetSubmissionDetail(submissionId int64, submission *problem_model.ProblemSubmissionDTO) error {

	err := global.DB.Table("problem_submissions AS s").
		Select("s.*, u.user_name").
		Joins("LEFT JOIN users u ON s.user_id = u.user_id").
		Where("s.id = ?", submissionId).
		First(submission).Error

	return err
}
