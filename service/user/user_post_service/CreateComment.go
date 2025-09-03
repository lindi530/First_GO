package user_post_service

import (
	"GO1/database/mysql/comment_mysql"
	"GO1/database/mysql/user_mysql"
	"GO1/models"
	"GO1/models/Comment"
	"GO1/service/comment_service"
	"time"
)

func CreateComment(userId, postId int64, requestComment *Comment.RequestComment) models.HandleFuncResp {
	comment := Comment.Comment{
		PostID:    postId,
		AuthorID:  userId,
		Content:   requestComment.Content,
		CreatedAt: time.Now(),
	}

	result := comment_mysql.CreateComment(&comment)

	if result.Ok {
		author := user_mysql.FindAuthorInfo(userId)
		result.Data = comment_service.BuildResponseComment(&comment, author, false)
	}

	return result
}
