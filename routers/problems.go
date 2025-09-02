package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func ProblemsRouter(router *gin.RouterGroup) {
	problems := router.Group("/problems")
	problems.GET("", api.ApiGroups.ProblemAPI.GetProblemList)
	problems.GET("/:problem_id", api.ApiGroups.ProblemAPI.GetProblemDetail)
	problems.Use(jwt.JWTAuthMiddleware())
	{
		problems.POST("/:problem_id/submit", api.ApiGroups.ProblemAPI.SubmitCode)
		problems.POST("/:problem_id/submit/example", api.ApiGroups.ProblemAPI.SubmitExample)
		problems.GET("/:problem_id/submissions", api.ApiGroups.ProblemAPI.GetProblemSubmissionList)
		problems.GET("/submissions", api.ApiGroups.ProblemAPI.GetUserSubmissionList)
		problems.GET("/submissions/:submission_id", api.ApiGroups.ProblemAPI.GetSubmissionDetail)
	}
}
