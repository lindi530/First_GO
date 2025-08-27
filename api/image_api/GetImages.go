package image_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/upload_model"
	"github.com/gin-gonic/gin"
)

func (ImageAPI) GetImages(c *gin.Context) {
	list := upload_model.Image{}
	global.DB.Find(&list)
	response.OkWithData(list, c)
}
