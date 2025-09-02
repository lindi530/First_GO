package post

import (
	"GO1/global"
	"GO1/models/post"
)

func GetPostByPostId(postId int64, post *post.Post) error {
	err := global.DB.
		Preload("Author").
		Where("id = ?", postId).
		First(post).Error
	return err
}

func GetAllPost() ([]post.Post, error) {
	var posts []post.Post
	err := global.DB.
		Preload("Author").
		Where("status = 0").
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error

	return posts, err
}
