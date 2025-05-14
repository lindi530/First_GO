package jwt

import (
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

const CUserIdKey = "user_id"

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.FailWithCode(response.NeedLogin, c)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			response.FailWithCode(response.InvalidAccessToken, c)
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set(CUserIdKey, claims.UserId)
		c.Next()
	}
}
