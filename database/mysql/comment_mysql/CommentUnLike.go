package comment_mysql

import (
	"GO1/global"
	"GO1/models/Comment"
	"gorm.io/gorm"
)

func CommentUnLike(commentId int64) {
	global.DB.Model(&Comment.Comment{}).
		Where("id = ?", commentId).
		UpdateColumn("likes", gorm.Expr("likes + ?", -1))
}
