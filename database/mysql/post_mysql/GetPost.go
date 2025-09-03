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

func GetAllPost() ([]post_model.Post, error) {
	var posts []post_model.Post
	err := global.DB.
		Preload("Author").
		Where("status = 0").
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error

	return posts, err
}
