package post

import (
	database_comment "GO1/database/mysql/comment"
	"GO1/database/mysql/user_mysql"
	"GO1/models"
	"GO1/models/Comment"
	service_comment "GO1/service/Comment"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateComment(c *gin.Context, userId, postId int64, requestComment *Comment.RequestComment) models.HandleFuncResp {
	comment := Comment.Comment{
		PostID:    postId,
		AuthorID:  userId,
		Content:   requestComment.Content,
		CreatedAt: time.Now(),
	}

	result := database_comment.CreateComment(&comment)

	if result.Ok {
		author := user_mysql.FindAuthorInfo(userId)
		result.Data = service_comment.BuildResponseComment(userId, &comment, author)
	}

	return result
}
