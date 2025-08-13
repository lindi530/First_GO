package problem_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
)

func (ProblemAPI) GetProblemList(c *gin.Context) {
	//
	userID := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.GetProblemList(userID)

	response.OkWithData(resp.Data, c)
}
