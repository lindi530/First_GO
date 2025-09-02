package post

import (
	"GO1/global"
	"GO1/models/post"
)

func PostCheckByPostID(postId int64) bool {
	var exists bool
	err := global.DB.Model(&post.Post{}).
		Select("count(1) > 0").
		Where("id = ?", postId).
		Scan(&exists).Error
	return err == nil
}
