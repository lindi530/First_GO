package post

import (
	"GO1/global"
	"GO1/models/post"
)

func CheckUserIDByPostID(postId int64, userId int64) bool {
	var id int64
	err := global.DB.
		Model(&post.Post{}).
		Where("id = ?", postId).
		Pluck("user_id", &id).Error
	if err != nil {
		return false
	}
	return id == userId
}
