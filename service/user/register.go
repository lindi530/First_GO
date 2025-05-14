package user

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
)

func Register(register models.ParamRegister) bool {
	name := mysql.UserNameParam(register.Name)
	result := mysql.CheckUser(name)

	if result {
		return false
	}

	mysql.Register(register)
	return true
}
