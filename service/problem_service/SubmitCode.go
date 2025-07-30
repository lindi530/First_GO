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

	problem_mysql.GetProblemExamples(codeSubmission.ProblemID, &examples)

	message := ws_model.MessageWs{
		From:    userid,
		To:      userid,
		Content: "Pending",
		Type:    "submit_code",
	}
	ws_service.WsHub.CodeStateWs(&message)
	resp.Code = 0
	runResult := RunCode(userid, codeSubmission.Code, codeSubmission.Language, &examples, &message)

	message.Content = "Accepted"
	for _, result := range runResult {
		if !result.Passed {
			resp.Code = 1
			message.Content = "Wrong Answer"
			break
		}
	}
	ws_service.WsHub.CodeStateWs(&message)
	resp.Data = runResult
	return
}
