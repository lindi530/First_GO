package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func CheckUserIDByPostID(postId int64, userId int64) bool {
	var id int64
	err := global.DB.
		Model(&post_model.Post{}).
		Where("id = ?", postId).
		Pluck("user_id", &id).Error
	if err != nil {
		return false
	}
	return id == userId
}
