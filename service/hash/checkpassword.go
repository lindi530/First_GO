package hash

import (
	"GO1/global"
	"GO1/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(user models.ParamLogin) bool {
	var dbPassword struct {
		password string
	}
	username, password := user.Username, user.Password
	global.DB.Where("username = ?", username).Select("username").First(&dbPassword)

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword.password), []byte(password))
	if err != nil {
		return false // 密码不正确
	}

	return true // 登录成功
}
