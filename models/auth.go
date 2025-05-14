package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserId               int64
	UserName             string
	JTI                  string // jwt id
	jwt.RegisteredClaims        // 内嵌标准的声明
}
