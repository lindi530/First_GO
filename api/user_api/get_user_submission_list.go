package user_api

import (
	"GO1/middlewares/response"
	"GO1/service/user_service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (UserAPI) GetUserSubmissionList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "2"))
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	resp := user_service.GetUserSubmissionList(userId, page, pageSize)

	response.DataAndMessage(&resp, c)
}
