package problem_api

import (
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
)

func (ProblemAPI) SubmitExample(c *gin.Context) {
	var exampleSubmit problem_model.ExampleSubmit
	if err := c.ShouldBindJSON(&exampleSubmit); err != nil {
		response.FailWithMessage("解析信息失败", c)
		return
	}
	if ok := isSafeCode(exampleSubmit.Code, exampleSubmit.Language); !ok {
		response.FailWithMessage("不安全代码", c)
		return
	}

	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.SubmitExample(userid, exampleSubmit)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithData(resp.Data, c)
}
