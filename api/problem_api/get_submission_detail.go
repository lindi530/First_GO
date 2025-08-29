package problem_api

import (
	"GO1/middlewares/response"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) GetSubmissionDetail(c *gin.Context) {
	submissionIdStr := c.Param("submission_id")
	submissionId, err := strconv.ParseInt(submissionIdStr, 10, 64)
	if err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	resp := problem_service.GetSubmissionDetail(submissionId)

	response.DataAndMessage(&resp, c)
}
