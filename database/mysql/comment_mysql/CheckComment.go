package comment_mysql

import (
	"GO1/global"
	"GO1/models/Comment"
)

func CheckCommentById(commentId int64) Comment.Comment {
	comment := Comment.Comment{}
	global.DB.Where("id = ?", commentId).First(&comment)
	return comment
}
