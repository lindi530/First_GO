package comment_service

import (
	"GO1/database/mysql/comment_like"
	"GO1/models"
)

func CommentLike(userId, commentId int64) (result models.HandleFuncResp) {
	err := comment_like.CommentLike(userId, commentId)
	if err != nil {
		result.Ok = false
		result.Msg = "点赞失败"
	} else {
		result.Ok = true
		result.Msg = "点赞成功"
	}
	return result
}
