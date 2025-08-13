package problem_api

import (
	"GO1/middlewares/response"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) GetProblemDetail(c *gin.Context) {
	problemID, _ := strconv.ParseUint(c.Param("problem_id"), 10, 10)

	resp := problem_service.GetProblemDetails(uint(problemID))

	if resp.Code == 1 {
		response.FailWithMessage("获取题目失败", c)
		return
	}

	response.OkWithData(resp.Data, c)
}
