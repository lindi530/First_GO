package post

import (
	mysql_comment "GO1/database/mysql/comment"
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
	"GO1/models/Comment"
)

func GetComments(userId, postId int64) models.HandleFuncResp {
	var comments []Comment.Comment
	mysql_comment.GetComments(&comments, postId)
	author := mysql_user.FindAuthorInfo(userId)

	var responseComments []Comment.ResponseComment
	for _, comment := range comments {
		res := Comment.BuildResponseComment(&comment, author)
		responseComments = append(responseComments, res)
	}

	return models.HandleFuncResp{
		Data: responseComments,
		Ok:   true,
	}
}
