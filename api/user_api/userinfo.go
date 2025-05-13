package user_api

import (
	"GO1/middlewares/response"
	"GO1/service/user"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) UserInfo(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.FailWithMessage(response.BadRequest, c)
	}
	userInfo := user.FindUserById(userid)
	response.OkWithData(userInfo, c)
}
