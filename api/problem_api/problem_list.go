package problem_api

import (
	"GO1/middlewares/response"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
)

func (ProblemAPI) GetProblemList(c *gin.Context) {
	//

	resp := problem_service.GetProblemList()

	response.OkWithData(resp.Data, c)
}
