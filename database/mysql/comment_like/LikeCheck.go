package comment_like

import (
	"GO1/global"
	"GO1/models/comment_like_model"
	"fmt"
)

func CommentLikeCheck(userId, commentId int64) bool {
	var like comment_like_model.CommentLike
	err := global.DB.
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		First(&like).Error
	fmt.Println("CommentLike err:", err == nil)
	if err != nil {
		return false
	}
	return true
}
