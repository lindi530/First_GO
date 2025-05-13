package user

import (
	"GO1/global"
	"GO1/models"
)

func DeleteUser(userid int) bool {
	user := models.User{}
	result := global.DB.Where("id = ?", userid).Delete(&user)
	return result.RowsAffected > 0
}
