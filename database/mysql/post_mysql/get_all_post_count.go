package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func GetAllPostCount(count *int64) {
	global.DB.Model(&post_model.Post{}).Count(count)
}
