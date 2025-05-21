package image_api

import (
	"GO1/middlewares/response"
	models_upload "GO1/models/upload"
	service "GO1/service/upload/images"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImageAPI) UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("无效数据1", c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		response.FailWithMessage("无效数据2", c)
		return
	}
	rsp := []models_upload.ResponseUploadImages{}
	service.UploadImage(c, fileList, &rsp)
	fmt.Println(rsp)
	response.OkWithData(rsp, c)
}
