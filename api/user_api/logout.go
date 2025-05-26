package user_api

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// 登出：从 Redis 删除对应 jti
func (UserAPI) Logout(c *gin.Context) {
	// 前端提交当前 refresh_token
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithCode(response.BadRequest, c)
	}
	claims, err := jwt.ParseToken(req.RefreshToken)
	if err != nil {
		response.FailWithCode(response.InvalidRefreshToken, c)
	}
	jti := claims.JTI

	redis.DeleteJWTId(c, jti)
	redis.DeleteOnlineState(c, claims.UserId)

	response.FailWithCode(response.Logout, c)
}
