package problem_api

import (
	"GO1/middlewares/response"
	"GO1/service/problem"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) GetProblemDetail(c *gin.Context) {
	problemID, _ := strconv.ParseInt(c.Param("problemID"), 10, 64)

	resp := problem.GetProblemDetails(problemID)

	if resp.Code == 1 {
		response.FailWithMessage("获取题目失败", c)
		return
	}

	response.OkWithData(resp.Data, c)
}
