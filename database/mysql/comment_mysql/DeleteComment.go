package comment_mysql

import (
	"GO1/global"
	"GO1/models/Comment"
)

func DeleteComment(comment Comment.Comment) {
	global.DB.Delete(&comment)
}
