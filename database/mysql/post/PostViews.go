package post

import (
	"GO1/global"
	"GO1/models/post"
	"gorm.io/gorm"
)

func PostViews(postId int64) {
	global.DB.
		Model(&post.Post{}).
		Where("id = ?", postId).
		UpdateColumns(map[string]interface{}{
			"views":      gorm.Expr("views + ?", 1),
			"updated_at": gorm.Expr("updated_at"),
		})
}
