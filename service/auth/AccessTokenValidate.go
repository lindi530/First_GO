package auth

import (
	"GO1/models"
	"GO1/pkg/jwt"
	"strings"
)

func AccessTokenValidate(userId int64, authHeader string) bool {
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := jwt.ParseToken(accessToken)
	if err != nil {
		return false
	}
	return claims.TokenType == models.AccessTokenType && claims.UserId == userId
}
