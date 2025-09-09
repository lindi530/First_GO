package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/service/user_service"
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
	result := user_service.ModifyAvatar(c, userId, avatar)

	global.Logger.Info(result)

	response.OkWithData(result, c)
}
