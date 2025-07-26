package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func ProblemsRouter(router *gin.RouterGroup) {
	problems := router.Group("/problems")
	problems.GET("/:problemID", api.ApiGroups.ProblemAPI.GetProblemDetail)
}
