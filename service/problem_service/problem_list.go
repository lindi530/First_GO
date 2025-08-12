package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
)

func GetProblemList() (resp response.Response) {
	problems, err := problem_mysql.GetProblemList()
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return
	}

	resp.Code = 0
	resp.Message = "成功查询题目列表"
	resp.Data = problems

	return
}
