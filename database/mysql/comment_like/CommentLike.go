package comment_like

import (
	"GO1/database/mysql/comment_mysql"
	"GO1/global"
	"GO1/models/comment_like_model"
)

func CommentLike(userId, commentId int64) error {
	like := comment_like_model.CommentLike{
		UserID:    userId,
		CommentID: commentId,
	}
	result := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		FirstOrCreate(&like).Error

	if result == nil {
		comment_mysql.CommentLike(commentId)
	}
	return result
}
