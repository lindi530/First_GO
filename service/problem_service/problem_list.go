package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
)

func GetProblemList(userID int64) (resp response.Response) {
	problems, err := problem_mysql.GetProblemList()
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return
	}

	if userID != 0 {
		userAcProblemIDs := problem_mysql.GetUserAcProblemID(userID)
		for _, problemID := range userAcProblemIDs {
			problems[problemID-1].Accepted = true
		}
	}

	resp.Code = 0
	resp.Message = "成功查询题目列表"
	resp.Data = problems

	return
}
