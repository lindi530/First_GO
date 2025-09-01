package comment_service

import (
	database_comment_like "GO1/database/mysql/comment_like"
	"GO1/models"
)

func CommentUnLike(userId, commentId int64) (result models.HandleFuncResp) {
	err := database_comment_like.CommentUnLike(userId, commentId)
	if err != nil {
		result.Ok = false
		result.Msg = "取消点赞失败"
	} else {
		result.Ok = true
		result.Msg = "取消成功"
	}
	return result
}
