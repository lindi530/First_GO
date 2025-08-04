package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func ProblemsRouter(router *gin.RouterGroup) {
	problems := router.Group("/problems")
	problems.GET("/:problemID", api.ApiGroups.ProblemAPI.GetProblemDetail)

	problems.Use(jwt.JWTAuthMiddleware())
	{
		problems.POST("/submit", api.ApiGroups.ProblemAPI.SubmitCode)
		problems.POST("/submit/example", api.ApiGroups.ProblemAPI.SubmitExample)
	}
}
