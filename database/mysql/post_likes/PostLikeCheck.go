package post_likes

import (
	"GO1/global"
	"GO1/models/post"
)

func PostLikeCheck(userId, postId int64) bool {
	post := post.PostLike{}
	err := global.DB.Model(&post).
		Where("user_id=? AND post_id=?", userId, postId).
		First(&post).Error
	if err != nil {
		return false
	}
	return true
}
