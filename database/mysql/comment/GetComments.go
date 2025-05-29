package comment

import (
	"GO1/global"
	"GO1/models/Comment"
)

func GetComments(comments *[]Comment.Comment, postId int64) {
	global.DB.Where("post_id = ?", postId).Order("created_at asc").Find(comments)
}
