package user

import (
	"GO1/global"
	"GO1/models"
)

func DeleteUser(userId int64) bool {
	user := models.User{}
	result := global.DB.Where("id = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}
