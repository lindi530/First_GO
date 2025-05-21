package image_api

import (
	"GO1/middlewares/response"
	models "GO1/models/Upload"
	service "GO1/service/Upload/images"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImageAPI) UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("无效数据", c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		response.FailWithMessage("无效数据", c)
		return
	}
	rsp := []*models.ResponseUploadImages{}
	service.UploadImage(c, fileList, &rsp)
	fmt.Println(rsp)
	response.OkWithData(rsp, c)
}
