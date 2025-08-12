package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
	"fmt"
)

func ModifyAvatar(userId int64, avatarMD5 string) {
	fmt.Println("ModifyAvatar", userId, avatarMD5)
	if CheckUser(UserIdParam(userId)) {
		global.DB.Model(user_model.User{}).
			Where("user_id = ?", userId).
			Update("avatar", avatarMD5)
	}
}
