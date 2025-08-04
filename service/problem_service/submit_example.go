package problem_service

import (
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"time"
)

func SubmitExample(userid int64, exampleSubmit problem_model.ExampleSubmit) (resp response.Response) {
	messageWs := ws_model.MessageWs{
		Id:       0,
		From:     userid,
		To:       userid,
		Content:  "Pending",
		Type:     "submit_code",
		State:    false,
		SendTime: time.Time{},
	}

	resp.Code = 0
	runResult := RunExample(exampleSubmit.Code, exampleSubmit.Language, exampleSubmit.Input, &messageWs)

	resp.Data = runResult
	return resp
}
