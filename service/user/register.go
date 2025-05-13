package user

import (
	database "GO1/database/user"
	"GO1/models"
)

func Register(register models.ParamRegister) bool {
	name := database.Name(register.Name)

	result := database.CheckUser(name)

	if result {
		return false
	}

	database.Register(register)
	return true
}
