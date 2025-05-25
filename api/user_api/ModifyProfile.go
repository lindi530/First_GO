package user_api

import (
	"GO1/middlewares/response"
	"GO1/models"
	service "GO1/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) ModifyProfile(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	var profile models.UserProfile
	if err := c.ShouldBindJSON(&profile); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	result := service.ModifyProfile(c, userId, profile)

	response.OkWithData(result, c)
}
