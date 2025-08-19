package match_api

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/match_service"
	"github.com/gin-gonic/gin"
)

func (MatchAPI) CancelMatch(c *gin.Context) {

	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := match_service.CancelMatch(userid)

	response.OkWithMessage(resp.Message, c)
}
