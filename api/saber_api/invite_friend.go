package saber_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/saber_service"
	"github.com/gin-gonic/gin"
)

func (SaberAPI) InviteFriend(c *gin.Context) {
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := saber_service.InviteFriend(userid)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithData(resp.Data, c)
}
