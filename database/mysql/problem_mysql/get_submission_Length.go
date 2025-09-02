package problem_mysql

import (
	"GO1/global"
	"GO1/models/problem_submission_model"
)

func GetUserSubmissionListLength(userId int64) (int64, error) {
	var length int64
	err := global.DB.
		Model(&problem_submission_model.ProblemSubmission{}).
		Where("user_id = ?", userId).Count(&length).Error

	return length, err
}
