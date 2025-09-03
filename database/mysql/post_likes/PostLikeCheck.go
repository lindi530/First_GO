package post_likes

import (
	"GO1/global"
	"GO1/models/post_model"
)

func PostLikeCheck(userId, postId int64) bool {
	var count int64
	// 只查询记录数量，性能更优
	err := global.DB.Model(&post_model.PostLike{}).
		Where("user_id = ? AND post_id = ?", userId, postId).
		Count(&count).Error

	if err != nil {
		return false
	}
	// 存在记录则已点赞
	return count > 0
}
