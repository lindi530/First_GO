package problem_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) GetUserSubmissionList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "2"))
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.GetUserSubmissionList(userId, page, pageSize)

	response.DataAndMessage(&resp, c)
}
