package comment_like

import (
	"GO1/database/mysql/comment_mysql"
	"GO1/global"
	"GO1/models/comment_like_model"
)

func CommentUnLike(userId, commentId int64) error {
	result := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		Delete(&comment_like_model.CommentLike{}).Error
	if result == nil {
		comment_mysql.CommentUnLike(commentId)
	}
	return result
}
