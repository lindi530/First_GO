package saber_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/saber_service"
	"github.com/gin-gonic/gin"
)

func (SaberAPI) JoinRoom(c *gin.Context) {
	roomId := c.Param("room_id")
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	resp := saber_service.JoinRoom(roomId, userId)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithMessage(resp.Message, c)
}
