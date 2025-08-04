package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"time"
)

func SubmitExample(userid, problemId int64, exampleSubmit problem_model.ExampleSubmit) (resp response.Response) {
	messageWs := ws_model.MessageWs{
		Id:       0,
		From:     userid,
		To:       userid,
		Content:  "Pending",
		Type:     "submit_code",
		State:    false,
		SendTime: time.Time{},
	}
	var constraints problem_model.ProblemConstraint
	err := problem_mysql.GetProblemConstraints(problemId, exampleSubmit.Language, &constraints)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return
	}

	resp.Code = 0
	runResult := RunExample(exampleSubmit.Code, exampleSubmit.Language,
		exampleSubmit.Input, constraints.MemoryLimit, constraints.TimeLimit, &messageWs)
	global.Logger.Info("constraints: ", constraints)
	resp.Data = runResult
	return resp
}
