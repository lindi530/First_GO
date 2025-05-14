package user

import (
	"GO1/pkg/jwt"
)

func Login(userid int64, username string) (string, string, string) {
	accessToken, _ := jwt.GenerateAccessToken(userid, username)
	refreshToken, jti, _ := jwt.GenerateRefreshToken(userid, username)

	return accessToken, refreshToken, jti
}
