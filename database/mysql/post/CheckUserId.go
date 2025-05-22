package post

import (
	"GO1/global"
	"GO1/models"
)

func CheckUserIDByPostID(postId int64, userId int64) bool {
	var id int64
	err := global.DB.
		Model(&models.Post{}).
		Where("id = ?", postId).
		Pluck("user_id", &id).Error
	if err != nil {
		return false
	}
	return id == userId
}
