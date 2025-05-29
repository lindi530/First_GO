package post

import (
	"GO1/global"
	"GO1/models/post"
)

func GetPostByPostId(postId int64) post.Post {
	post := post.Post{}
	global.DB.Where("id = ?", postId).First(&post)
	return post
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
