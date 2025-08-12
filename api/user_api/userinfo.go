package user_api

import (
	mysql "GO1/database/mysql/user_mysql"
	"GO1/middlewares/response"
	"GO1/models/user_model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) UserInfo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		response.FailWithMessage(response.BadRequest, c)
	}
	userInfo := mysql.FindUser(mysql.UserIdParam(userId))
	fmt.Println(userInfo)

	responseUser := user_model.BuildUserResponse(c, userInfo)

	response.OkWithData(responseUser, c)
}
