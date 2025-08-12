package match_api

import (
	"GO1/pkg/jwt"
	"GO1/service/match_service"
	"github.com/gin-gonic/gin"
)

func (MatchAPI) MatchRequest(c *gin.Context) {
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	err := match_service.SendMatchRequest(userid)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"status": "queued for match_model"})
	}
}
