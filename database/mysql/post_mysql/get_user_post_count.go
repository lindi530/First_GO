package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func GetUserPostCount(userId int64, count *int64) error {
	return global.DB.
		Model(post_model.Post{}).
		Where("user_id=?", userId).
		Count(count).Error
}
