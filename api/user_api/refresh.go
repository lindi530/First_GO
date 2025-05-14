package user_api

import (
	"GO1/database/redis"
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (UserAPI) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithCode(response.BadRequest, c)
		return
	}
	claims, err := jwt.ParseToken(req.RefreshToken)
	if err != nil {
		response.FailWithCode(response.InvalidRefreshToken, c)
		return
	}
	jti, userID, userName := claims.JTI, claims.UserId, claims.UserName
	// 检查 Redis 中是否存在此 jti
	if result := redis.CheckRefreshToken(c, userID, jti); !result {
		response.FailWithCode(response.InvalidRefreshToken, c)
		return
	}
	// 2) 删除旧 jti （一次性使用）
	//redis.DeleteJWTId(c, jti)

	// 3) 签发新的令牌对           只签发accessToken
	accessToken, _ := jwt.GenerateAccessToken(userID, userName)
	//refreshToken, newJti, _ := jwt.GenerateRefreshToken(userID, userName)

	//redis.SaveJWTId(c, userID, newJti)

	c.JSON(http.StatusOK, gin.H{
		"username":     userName,
		"access_token": accessToken,
	})
}
