package routers

import (
	"GO1/api"
	"GO1/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func UploadRouter(router *gin.RouterGroup) {
	upload := router.Group("/upload")

	upload.Use(jwt.JWTAuthMiddleware())
	{
		upload.POST("problem", api.ApiGroups.UploadAPI.UploadProblem)
	}
}
