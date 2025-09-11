package user_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_submission_model"
	"github.com/gin-gonic/gin"
)

func GetUserSubmissionList(userId int64, page, pageSize int) (resp response.Response) {
	var submissions []problem_submission_model.UserSubmission

	offset := (page - 1) * pageSize
	limit := pageSize
	if err := problem_mysql.GetUserSubmissionList(&submissions, userId, offset, limit); err != nil {
		resp.Code = 1
		resp.Message = "数据获取失败"
		return
	}

	length, err := problem_mysql.GetUserSubmissionListLength(userId)
	if err != nil {
		resp.Code = 1
		resp.Message = "数据获取失败"
		return
	}
	resp.Code = 0
	resp.Data = gin.H{
		"list":  submissions,
		"total": length,
	}
	return
}
