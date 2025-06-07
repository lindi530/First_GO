package chat_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/ws"
	"github.com/gin-gonic/gin"
)

func (ChatAPI) ChatHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.FailWithCode(response.InvalidAccessToken, c)
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		response.FailWithCode(response.InvalidAccessToken, c)
		return
	}

	ws.ConstructWS(c, claims.UserId)
}
