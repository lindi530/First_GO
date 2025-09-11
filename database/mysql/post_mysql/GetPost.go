package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func GetPostByPostId(postId int64, post *post_model.Post) error {
	err := global.DB.
		Preload("Author").
		Where("id = ?", postId).
		First(&post).
		Error
	return err
}

func GetAllPost(posts *[]post_model.Post, offset, limit int) error {

	limit = min(limit, 10)
	err := global.DB.
		Preload("Author").
		Where("status = 0").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error

	return err
}
