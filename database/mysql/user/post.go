package user

import (
	"GO1/global"
	"GO1/models/post"
)

func GetUserPost(userID int64) ([]post.Post, error) {
	var posts []post.Post
	err := global.DB.
		Preload("Author").
		Where("user_id = ? AND status = 0", userID).
		Order("created_at DESC").
		Find(&posts).Error
	return posts, err
}

func CreatePost(post *post.Post) {
	global.DB.Create(post) // create 固定执行插入操作
	//global.DB.Save(&post)     // save 先判断主键，主键存在则修改该信息，不存在进行插入操作
}
