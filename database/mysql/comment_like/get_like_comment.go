package comment_like

import (
	"GO1/global"
	"GO1/models/comment_like_model"
)

func GetUserLikeCommentIDs(likeIDs, commentIDs *[]int64, userid int64) error {
	err := global.DB.Model(&comment_like_model.CommentLike{}).
		Select("comment_id").
		Where("user_id = ? AND comment_id IN ?", userid, *commentIDs).
		Scan(likeIDs).Error
	return err
}
