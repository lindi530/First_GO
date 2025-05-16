package user

import (
	"GO1/global"
	"GO1/models"
)

func GetPost(userID int64) ([]models.Post, error) {
	var posts []models.Post
	err := global.DB.
		Preload("Author").
		Where("user_id = ? AND status = 0", userID).
		Order("created_at DESC").
		Find(&posts).Error
	return posts, err
}
