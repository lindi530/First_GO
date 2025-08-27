package upload_service

import (
	"GO1/database/mysql/problem_mysql"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
)

func UploadProblem(data *problem_model.UploadProblem) (resp response.Response) {

	problem := data.MapProblem()

	err := problem_mysql.UploadProblem(problem)
	if err != nil {
		resp.Code = 1
		resp.Message = "题目保存失败！"
		return
	}

	resp.Code = 0
	resp.Message = "题目保存成功！"

	return
}
