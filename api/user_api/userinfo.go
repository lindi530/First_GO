package user_api

import (
	"GO1/database/mysql/user_mysql"
	"GO1/middlewares/response"
	"GO1/models/user_model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) UserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		response.FailWithMessage(response.BadRequest, c)
	}
	userInfo := user_mysql.FindUser(user_mysql.UserIdParam(userId))

	responseUser := user_model.BuildUserResponse(c, userInfo)

	response.OkWithData(responseUser, c)
}
