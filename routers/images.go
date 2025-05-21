package routers

import (
	"GO1/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ImagesRouter(router *gin.RouterGroup) {
	images := router.Group("/images")
	fmt.Println(images.BasePath())
	images.POST("", api.ApiGroups.ImageAPI.UploadImages)
}
