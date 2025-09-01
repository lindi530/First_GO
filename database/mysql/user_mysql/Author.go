package user_mysql

import (
	"GO1/global"
	"GO1/models/user_model"
)

func FindAuthorInfo(userId int64) *user_model.AuthorInfo {
	var author user_model.AuthorInfo
	err := global.DB.
		Model(&user_model.User{}).
		Select("user_id", "user_name", "avatar").
		Where("user_id = ?", userId).
		First(&author).Error
	if err != nil {
		return nil
	}
	return &author
}

func FindAuthorsInfo(authors *[]user_model.AuthorInfo, authorIDs []int64) error {
	err := global.DB.
		Model(&user_model.User{}).
		Select("user_id", "user_name", "avatar").
		Where("user_id in ?", authorIDs).
		Find(&authors).Error
	return err
}
