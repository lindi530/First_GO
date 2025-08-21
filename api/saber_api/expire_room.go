package saber_api

import (
	"GO1/middlewares/response"
	"GO1/service/saber_service"
	"github.com/gin-gonic/gin"
)

func (SaberAPI) ExpireRoom(c *gin.Context) {
	roomId := c.Param("room_id")

	resp := saber_service.ExpireRoom(roomId)

	if resp.Code == 1 {
		response.FailWithMessage("房间删除失败", c)
		return
	}

	response.OkWithMessage("房间删除成功", c)

}
