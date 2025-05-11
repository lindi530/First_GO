package user

import (
	database "GO1/database/user"
	"GO1/models"
)

func Register(user models.Register) bool {
	name := database.Name(user.Name)

	result := database.CheckUser(name)

	if result {
		return false
	}
	database.Register(user)
	return true
}
