package saber_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/saber_model"
	"GO1/pkg/jwt"
	"GO1/service/saber_service"
	"github.com/gin-gonic/gin"
)

func (SaberAPI) SaberSubmit(c *gin.Context) {
	var data saber_model.SaberSubmit
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWithMessage("提交数据格式失败", c)
		return
	}
	global.Logger.Error("roomId ", data)
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := saber_service.SaberSubmit(userid, &data)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithMessage(resp.Message, c)
}
