package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
)

func DeletePost(post post_model.Post) {
	global.DB.Delete(post)
}
