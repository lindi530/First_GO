package auth_api

import (
	"GO1/middlewares/response"
	"GO1/service/auth"
	"github.com/gin-gonic/gin"
)

func (AuthAPI) RefreshTokenValidate(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}

	result := auth.RefreshTokenValidate(c, req.RefreshToken)
	if !result.Ok {
		response.FailWithCode(response.InvalidRefreshToken, c)
		return
	}

	response.OkWithData(result.Data, c)
}
