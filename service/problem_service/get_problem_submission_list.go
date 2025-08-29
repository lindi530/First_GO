package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
)

func GetProblemSubmissionList(userId int64, problemId uint) (resp response.Response) {

	var submissions []problem_model.ProblemSubmission

	if err := problem_mysql.GetProblemSubmissionList(userId, problemId, &submissions); err != nil {
		resp.Code = 1
		resp.Message = "提交记录查询失败"
		return
	}

	resp.Code = 0
	resp.Data = submissions
	resp.Message = "提交记录查询成功"

	return
}
