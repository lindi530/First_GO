package problem

import (
	mysql "GO1/database/mysql/problem"
	"GO1/middlewares/response"
	models "GO1/models/problem"
)

func GetProblemDetails(id int64) (resp response.Response) {
	var problemDetail models.Problem

	if err := mysql.GetProblemDetails(id, &problemDetail); err != nil {
		resp.Code = 1
		resp.Message = "读取题目信息失败"
		return
	}

	if err := mysql.GetProblemTags(id, &problemDetail); err != nil {
		resp.Code = 1
		resp.Message = "读取题目信息失败"
		return
	}

	if err := mysql.GetProblemExamples(id, &problemDetail); err != nil {
		resp.Code = 1
		resp.Message = "读取题目信息失败"
		return
	}

	resp.Code = 0
	resp.Data = problemDetail
	resp.Message = "读取题目信息成功"

	return
}
