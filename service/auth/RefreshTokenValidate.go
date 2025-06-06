package auth

import (
	"GO1/database/redis"
	"GO1/models"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func RefreshTokenValidate(c *gin.Context, RefreshToken string) (resp models.HandleFuncResp) {
	claims, err := jwt.ParseToken(RefreshToken)
	if err != nil {
		resp.Ok = false
		return
	}
	jti, userID, userName := claims.JTI, claims.UserId, claims.UserName
	// 检查 Redis 中是否存在此 jti
	if result := redis.CheckRefreshToken(c, userID, jti); !result {
		resp.Ok = false
		return
	}
	//2) 删除旧 jti （一次性使用）
	redis.DeleteJWTId(c, jti)

	// 3) 签发新的令牌对           只签发accessToken
	accessToken, _ := jwt.GenerateAccessToken(userID, userName)
	//refreshToken, newJti, _ := jwt.GenerateRefreshToken(userID, userName)

	//redis.SaveJWTId(c, userID, newJti)
	resp.Ok = true
	resp.Data = gin.H{
		"access_token": accessToken,
	}
	return
}
