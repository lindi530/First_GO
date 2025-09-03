package post_mysql

import (
	"GO1/global"
	"GO1/models/post_model"
	"gorm.io/gorm"
)

func PostViews(postId int64) {
	global.DB.
		Model(&post_model.Post{}).
		Where("id = ?", postId).
		UpdateColumns(map[string]interface{}{
			"views":      gorm.Expr("views + ?", 1),
			"updated_at": gorm.Expr("updated_at"),
		})
}
