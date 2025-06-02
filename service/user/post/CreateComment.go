package post

import (
	database_comment "GO1/database/mysql/comment"
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
	"GO1/models/Comment"
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
		author := mysql_user.FindAuthorInfo(userId)
		result.Data = Comment.BuildResponseComment(c, userId, &comment, author)
	}

	return result
}
