package image_api

import (
	"GO1/middlewares/response"
	"GO1/models"
	"GO1/models/upload_model"
	"GO1/service"
	"github.com/gin-gonic/gin"
)

func (ImageAPI) DeleteImages(c *gin.Context) {
	var deleteList models.DeleteRequest
	if err := c.ShouldBindJSON(&deleteList); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	err := service.DeleteList(deleteList.IdList, upload_model.Image{})
	if err != nil {
		response.FailWithCode(response.DeleteFail, c)
	}

	response.OkWithData("删除成功", c)
}
