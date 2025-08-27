package upload_api

import (
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/service/upload_service"
	"github.com/gin-gonic/gin"
)

func (UploadAPI) UploadProblem(c *gin.Context) {

	problemData := problem_model.UploadProblem{}
	if err := c.ShouldBindJSON(&problemData); err != nil {
		response.FailWithMessage("上传信息有误", c)
		return
	}

	resp := upload_service.UploadProblem(&problemData)

	response.MessageAndMessage(&resp, c)
}
