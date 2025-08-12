package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/user_model"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) ModifyProfile(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var profile user_model.UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		global.Logger.Error(err)

		response.FailWithCode(response.BadRequest, c)
		return
	}
	global.Logger.Info("service ModifyProfile")
	result := service.ModifyProfile(c, userId, profile)
	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
	}
	response.OkWithMessage(result.Msg, c)
}
