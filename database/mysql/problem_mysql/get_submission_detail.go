package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_submission_model"
)

func GetSubmissionDetail(submissionId int64, submission *problem_submission_model.ProblemSubmissionDTO) error {

	err := global.DB.Table("problem_submissions AS s").
		Select("s.*, u.user_name, p.title, p.description").
		Joins("LEFT JOIN users u ON s.user_id = u.user_id").
		Joins("LEFT JOIN problems p ON s.problem_id = p.id").
		Where("s.id = ?", submissionId).
		First(submission).Error

	return err
}
