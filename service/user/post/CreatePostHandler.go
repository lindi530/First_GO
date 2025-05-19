package post

import (
	"GO1/models"
)

func CreatePostHandler(p models.CreatePost) models.Post {
	// 转换为Post模型（忽略PostID，由数据库自动生成）
	post := models.Post{
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return post
}
