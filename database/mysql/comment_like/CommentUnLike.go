package comment_like

import (
	"GO1/global"
	models_CommentLike "GO1/models/CommentLike"
)

func CommentUnLike(userId, commentId int64) error {
	return global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		Delete(&models_CommentLike.CommentLike{}).Error
}
