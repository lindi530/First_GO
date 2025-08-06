package match_api

import (
	"GO1/middlewares/response"
	"GO1/models/match_model"
	"GO1/service/match_service"
	"github.com/gin-gonic/gin"
)

func (MatchAPI) GetMatch(c *gin.Context) {

	matchUser := match_model.MatchUser{}
	if err := c.ShouldBindJSON(&matchUser); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	err := match_service.SendMatchRequest(matchUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"status": "queued for match_model"})
	}
}
