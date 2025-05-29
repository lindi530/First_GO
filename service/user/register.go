package user

import (
	mysql "GO1/database/mysql/user"
	"GO1/models/user"
)

func Register(register user.ParamRegister) bool {
	result := mysql.CheckUser(mysql.UserNameParam(register.Name))

	if result {
		return false
	}
	mysql.Register(register)
	return true
}
