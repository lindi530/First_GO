package user

import (
	"GO1/global"
	"GO1/models"
)

func ModifyAvatar(userId int64, avatarMD5 string) {
	if CheckUser(UserIdParam(userId)) {
		global.DB.Model(models.User{}).
			Where("user_id = ?", userId).
			Update("avatar", avatarMD5)
	}
}
