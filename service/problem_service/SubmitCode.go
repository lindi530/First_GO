package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
)

func SubmitCode(userid int64, codeSubmission problem_model.CodeSubmission) (resp response.Response) {
	var examples []problem_model.Example

	err := problem_mysql.GetProblemExamples(codeSubmission.ProblemID, &examples)
	if err != nil {
		resp.Code = 1
		resp.Message = "判定失败，请重试"
		return
	}

	message := &ws_model.EditStatus{
		Type: ws_model.MessageTypeEditStatus,
		To:   userid,
	}
	ws_service.WsHub.SendEditData(message, "Pending")
	resp.Code = 0
	runResult := RunCode(userid, codeSubmission.Code, codeSubmission.Language, &examples, message)

	msgContent := "Accepted"
	for _, result := range runResult {
		if !result.Passed {
			resp.Code = 1
			msgContent = "Wrong Answer"
			break
		}
	}
	ws_service.WsHub.SendEditData(message, msgContent)
	resp.Data = runResult
	return
}
