package jwt

import (
	"GO1/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

// 生成 JWT
func GenerateAccessToken(userID int64, userName string) (string, error) {

	claims := models.CustomClaims{
		userID,
		userName, // 自定义字段
		models.AccessTokenType,
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
		models.RefreshTokenType,
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
func ParseToken(tokenStr string) (*models.CustomClaims, error) {
	claims := &models.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		// 直接用 errors.Is 判断是否为 token 过期错误
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}

func GetUserIdFromToken(authHeader string) int64 {
	claims, err := GetUserClaims(authHeader)
	if err != nil {
		return 0
	}
	return claims.UserId
}

func GetUserClaims(authHeader string) (*models.CustomClaims, error) {
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := ParseToken(accessToken)
	return claims, err
}
