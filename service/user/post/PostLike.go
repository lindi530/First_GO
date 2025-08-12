package post

import (
	"GO1/database/mysql/post"
	"GO1/database/mysql/post_likes"
	"GO1/database/mysql/user_mysql"
	"GO1/models"
)

func PostLike(userId, postId int64) (resp models.HandleFuncResp) {
	exit1 := user_mysql.CheckUserByUserId(userId)
	exit2 := post.PostCheckByPostID(postId)

	if !exit1 || !exit2 {
		resp.Msg = "点赞帖子失败，不存在该帖子或用户"
		resp.Ok = false
		return
	}

	result := post_likes.PostLike(userId, postId)

	return result
}
