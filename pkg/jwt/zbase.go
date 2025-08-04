package jwt

import (
	"GO1/global"
	"errors"
	"time"
)

var (
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
)

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("token is invalid")
)

var jwtSecret = []byte("生产队的代码驴")

func InitJwt() {
	AccessTokenDuration = time.Minute * time.Duration(global.Config.JWT.AccessDuration)
	RefreshTokenDuration = time.Hour * time.Duration(global.Config.JWT.RefreshDuration)
}
