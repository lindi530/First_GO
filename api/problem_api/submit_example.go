package problem_api

import (
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
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

	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	problemIDStr := c.Param("problem_id")

	problemID64, err := strconv.ParseInt(problemIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的题目ID", c)
		return
	}
	resp := problem_service.SubmitExample(userId, problemID64, exampleSubmit)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithData(resp.Data, c)
}
