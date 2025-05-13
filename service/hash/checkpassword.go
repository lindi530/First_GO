package hash

import (
	"GO1/global"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(username string, password string) bool {
	var dbPassword struct {
		password string
	}

	global.DB.Where("username = ?", username).Select("username").First(&dbPassword)

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword.password), []byte(password))
	if err != nil {
		return false // 密码不正确
	}

	return true // 登录成功
}
