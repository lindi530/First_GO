package message_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/messages"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (MessageAPI) GetMessageByTargetId(c *gin.Context) {
	targetId, _ := strconv.ParseInt(c.Param("target_id"), 10, 64)
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := messages.GetMessageByTargetId(targetId, userId)

	if !resp.Ok {
		response.FailWithCode(response.FindMessagesFail, c)
		return
	}

	response.OkWithData(resp.Data, c)
}
