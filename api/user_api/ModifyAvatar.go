package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) ModifyAvatar(c *gin.Context) {
	avatar, err := c.FormFile("avatar")
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}
	err = service.ModifyAvatar(c, userId, avatar)
	if err != nil {
		global.Logger.Error(err)
		response.FailWithCode(response.UploadFail, c)
		return
	}

	response.OkWithMessage("上传成功", c)
}
