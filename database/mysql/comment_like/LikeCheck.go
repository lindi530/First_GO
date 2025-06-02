package comment_like

import (
	"GO1/global"
	models_CommentLike "GO1/models/CommentLike"
	"fmt"
)

func CommentLikeCheck(userId, commentId int64) bool {
	var like models_CommentLike.CommentLike
	err := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		First(&like).Error
	fmt.Println("CommentLike err:", err == nil)
	if err != nil {
		return false
	}
	return true
}
