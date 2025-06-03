package post

import (
	"GO1/global"
	"GO1/models/post"
	"gorm.io/gorm"
)

func PostViews(postId int64) {
	global.DB.Model(&post.Post{}).
		Where("id = ?", postId).
		UpdateColumn("Views", gorm.Expr("views + ?", 1))
}
