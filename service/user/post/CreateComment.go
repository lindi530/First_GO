package post

import (
	database_comment "GO1/database/mysql/comment"
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
	"GO1/models/Comment"
	"time"
)

func CreateComment(userId, postId int64, requestComment Comment.RequestComment) models.HandleFuncResp {
	comment := Comment.Comment{
		PostID:    postId,
		AuthorID:  userId,
		Content:   requestComment.Content,
		CreatedAt: time.Now(),
	}

	result := database_comment.CreateComment(&comment)

	if result.Ok {
		user := mysql_user.FindAuthorInfo(userId)
		result.Data = Comment.BuildResponseComment(&comment, user)
	}

	return result
}
