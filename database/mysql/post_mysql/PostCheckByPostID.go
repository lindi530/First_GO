package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func PostCheckByPostID(postId int64) bool {
	var exists bool
	err := global.DB.Model(&post_model.Post{}).
		Select("count(1) > 0").
		Where("id = ?", postId).
		Scan(&exists).Error
	return err == nil
}
