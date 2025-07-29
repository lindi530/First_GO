package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func ProblemsRouter(router *gin.RouterGroup) {
	problems := router.Group("/problems")
	problems.GET("/:problemID", api.ApiGroups.ProblemAPI.GetProblemDetail)
	problems.POST("/submit", api.ApiGroups.ProblemAPI.SubmitCode)
	problems.Use(jwt.JWTAuthMiddleware())
	{

	}
}
