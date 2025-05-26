package user_api

import (
	"GO1/pkg/jwt"
	"strings"
)

type UserAPI struct{}

func GetUserIdFromToken(authHeader string) int64 {
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	claims, _ := jwt.ParseToken(accessToken)

	return claims.UserId
}
