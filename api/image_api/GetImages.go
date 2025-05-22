package image_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/upload"
	"github.com/gin-gonic/gin"
)

func (ImageAPI) GetImages(c *gin.Context) {
	list := upload.Image{}
	global.DB.Find(&list)
	response.OkWithData(list, c)
}
