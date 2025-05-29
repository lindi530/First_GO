package post

import (
	database_comment "GO1/database/mysql/comment"
	"GO1/models"
	"GO1/models/Comment"
)

func GetComments(postId int64) models.HandleFuncResp {
	var comments []Comment.Comment
	database_comment.GetComments(&comments, postId)

	return models.HandleFuncResp{
		Data: comments,
		Ok:   true,
	}
}
