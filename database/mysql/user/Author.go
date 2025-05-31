package user

import (
	"GO1/global"
	models_user "GO1/models/user"
)

func FindAuthorInfo(userId int64) *models_user.AuthorInfo {
	var author models_user.AuthorInfo
	err := global.DB.
		Model(&models_user.User{}).
		Select("user_id", "user_name", "avatar").
		Where("user_id = ?", userId).
		First(&author).Error
	if err != nil {
		return nil
	}
	return &author
}
