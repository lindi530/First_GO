package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
)

func DeleteUser(userId int64) bool {
	user := user_model.User{}
	result := global.DB.Where("id = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}
