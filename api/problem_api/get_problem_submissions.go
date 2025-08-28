package problem_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) GetProblemSubmissions(c *gin.Context) {
	problemIdStr := c.Param("problem_id")
	problemId, err := strconv.ParseUint(problemIdStr, 10, 64)
	if err != nil {
		response.FailWithCode(response.BadRequest, c)
	}
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.GetProblemSubmissions(userId, uint(problemId))

	response.DataAndMessage(&resp, c)
}
