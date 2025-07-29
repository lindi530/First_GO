package problem_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
)

func SubmitCode(userid int64, codeSubmission problem_model.CodeSubmission) (resp response.Response) {
	var examples []problem_model.Example

	problem_mysql.GetProblemExamples(codeSubmission.ProblemID, &examples)

	resp.Code = 0
	var runResult []problem_model.RunResult
	for _, example := range examples {
		res := RunCode(codeSubmission.Code, codeSubmission.Language, example.Input, example.Output)
		runResult = append(runResult, res)
		if !res.Passed {
			resp.Code = 1
		}
	}
	resp.Data = runResult
	return
}
