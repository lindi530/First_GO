package user_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func GetUserPost(userID int64, page, pageSize int, posts *[]post_model.Post) error {
	offset := (page - 1) * pageSize
	limit := pageSize
	err := global.DB.
		Preload("Author").
		Where("user_id = ? AND status = 0", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error
	return err
}

func CreatePost(post *post_model.Post) {
	global.DB.Create(post) // create 固定执行插入操作
	//global.DB.Save(&post)     // save 先判断主键，主键存在则修改该信息，不存在进行插入操作
}
