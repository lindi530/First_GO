package auth

import (
	"GO1/models"
	"GO1/pkg/jwt"
	"strings"
)

func RefreshTokenValidate(userId int64, authHeader string) bool {
	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		return false
	}
	return claims.TokenType == models.RefreshTokenType && claims.UserId == userId
}
