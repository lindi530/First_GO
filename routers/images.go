package routers

import (
	"GO1/api"
	"github.com/gin-gonic/gin"
)

func ImagesRouter(router *gin.RouterGroup) {
	images := router.Group("/images")
	images.GET("", api.ApiGroups.ImageAPI.GetImages)
	images.POST("", api.ApiGroups.ImageAPI.UploadImages)
	images.DELETE("", api.ApiGroups.ImageAPI.DeleteImages)
}
