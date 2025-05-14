package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(dbPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return false // 密码不正确
	}

	return true // 登录成功
}
