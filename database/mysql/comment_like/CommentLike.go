package comment_like

import (
	"GO1/global"
	model_CommentLike "GO1/models/CommentLike"
)

func CommentLike(userId, commentId int64) error {
	like := model_CommentLike.CommentLike{
		UserID:    userId,
		CommentID: commentId,
	}
	result := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		FirstOrCreate(&like).Error

	return result
}
