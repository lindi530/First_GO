package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
)

func GetSubmissionDetail(submissionId int64) (resp response.Response) {

	var submissionDetail problem_model.ProblemSubmissionDTO
	if err := problem_mysql.GetSubmissionDetail(submissionId, &submissionDetail); err != nil {
		resp.Code = 1
		resp.Message = "查询详细记录失败！"
		return
	}

	resp.Code = 0
	resp.Message = "查询记录成功"
	resp.Data = submissionDetail
	return
}
