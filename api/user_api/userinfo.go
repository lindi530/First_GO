package user_api

import (
	mysql "GO1/database/mysql/user"
	"GO1/middlewares/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) UserInfo(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage(response.BadRequest, c)
	}
	userInfo := mysql.FindUser(mysql.UserIdParam(userid))

	response.OkWithData(userInfo, c)
}
