package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
)

func SubmitExample(userid, problemId int64, exampleSubmit problem_model.ExampleSubmit) (resp response.Response) {
	messageWs := &ws_model.MessageWs{
		Type: ws_model.MessageTypeEditStatus,
		To:   userid,
	}
	ws_service.WsHub.CodeStateWs(messageWs, "Pending")
	var constraints problem_model.ProblemConstraint
	err := problem_mysql.GetProblemConstraints(problemId, exampleSubmit.Language, &constraints)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return
	}

	resp.Code = 0
	runResult := RunExample(exampleSubmit.Code, exampleSubmit.Language,
		exampleSubmit.Input, constraints.MemoryLimit, constraints.TimeLimit, messageWs)
	global.Logger.Info("constraints: ", constraints)
	ws_service.WsHub.CodeStateWs(messageWs, runResult.Error)
	resp.Data = runResult
	return resp
}
