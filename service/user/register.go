package user

import (
	"GO1/database/mysql/user_mysql"
	"GO1/models/user_model"
)

func Register(register user_model.ParamRegister) bool {
	result := user_mysql.CheckUser(user_mysql.UserNameParam(register.Name))

	if result {
		return false
	}
	user_mysql.Register(register)
	return true
}
