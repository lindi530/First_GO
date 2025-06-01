package user

import (
	"GO1/global"
	"GO1/models/user"
	"fmt"
)

func ModifyAvatar(userId int64, avatarMD5 string) {
	fmt.Println("ModifyAvatar", userId, avatarMD5)
	if CheckUser(UserIdParam(userId)) {
		global.DB.Model(user.User{}).
			Where("user_id = ?", userId).
			Update("avatar", avatarMD5)
	}
}
