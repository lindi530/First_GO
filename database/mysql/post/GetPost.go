package post

import (
	"GO1/global"
	"GO1/models"
)

func GetPostByPostId(postId int64) models.Post {
	post := models.Post{}
	global.DB.Where("post_id = ?", postId).First(&post)
	return post
}

func GetAllPost() ([]models.Post, error) {
	var posts []models.Post
	err := global.DB.
		Preload("Author").
		Where("status = 0").
		Order("created_at DESC").
		Limit(10).
		Find(&posts).Error
	return posts, err
}
