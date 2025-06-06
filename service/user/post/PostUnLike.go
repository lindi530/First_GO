package post

import (
	"GO1/database/mysql/post_likes"
	"GO1/models"
)

func PostUnLike(userId, postId int64) (resp models.HandleFuncResp) {
	exit := post_likes.PostLikeCheck(userId, postId)

	if !exit {
		resp.Msg = "该点赞关系不存在"
		resp.Ok = false
		return
	}

	result := post_likes.PostUnLike(userId, postId)

	return result
}
