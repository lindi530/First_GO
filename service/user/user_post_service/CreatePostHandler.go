package user_post_service

import (
	"GO1/models/post_model"
	"time"
)

func CreatePostHandler(p post_model.CreatePost) post_model.Post {
	// 转换为Post模型（忽略PostID，由数据库自动生成）
	post := post_model.Post{
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Status:    p.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return post
}
