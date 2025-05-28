package jwt

import (
	"GO1/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

const (
	AccessTokenDuration  = time.Hour * 1
	RefreshTokenDuration = time.Hour * 8
)

var jwtSecret = []byte("生产队的代码驴")

// 生成 JWT
func GenerateAccessToken(userID int64, userName string) (string, error) {
	claims := models.CustomClaims{
		userID,
		userName, // 自定义字段
		"",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
			Issuer:    "GO1", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID int64, userName string) (string, string, error) {
	jti := uuid.NewString() // 唯一标识
	claims := models.CustomClaims{
		userID,
		userName,
		jti,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)), // 过期时间
			Issuer:    "GO1",                                                    // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString(jwtSecret)
	return s, jti, err
}

// 解析 JWT
func ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GetUserIdFromToken(authHeader string) int64 {
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	claims, _ := ParseToken(accessToken)

	return claims.UserId
}
