package user_api

import (
	"GO1/middlewares/response"
	"GO1/models/user"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) ModifyProfile(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var profile user.UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	result := service.ModifyProfile(c, userId, profile)
	if !result.Ok {
		response.FailWithMessage(result.Msg, c)
	}
	response.OkWithMessage(result.Msg, c)
}
