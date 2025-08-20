package problem_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ProblemAPI) SubmitCode(c *gin.Context) {
	var codeSubmission problem_model.CodeSubmission
	if err := c.ShouldBindJSON(&codeSubmission); err != nil {
		response.FailWithMessage("解析信息失败", c)
		return
	}

	problemIDStr := c.Param("problem_id") // 返回 string
	problemID, err := strconv.ParseUint(problemIDStr, 10, 64)
	if err != nil {
		response.FailWithMessage("解析信息失败", c)
		return
	}
	
	codeSubmission.ProblemID = uint(problemID)
	global.Logger.Info("SubmitCode:", codeSubmission.ProblemID)
	if ok := isSafeCode(codeSubmission.Code, codeSubmission.Language); !ok {
		response.FailWithMessage("不安全代码", c)
		return
	}
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.SubmitCode(userid, codeSubmission)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithData(resp.Data, c)
}
