package comment_like

import (
	"GO1/database/mysql/comment"
	"GO1/global"
	models_CommentLike "GO1/models/CommentLike"
)

func CommentUnLike(userId, commentId int64) error {
	result := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		Delete(&models_CommentLike.CommentLike{}).Error
	if result == nil {
		comment.CommentUnLike(commentId)
	}
	return result
}
