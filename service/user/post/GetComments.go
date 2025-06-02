package post

import (
	mysql_comment "GO1/database/mysql/comment"
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
	"GO1/models/Comment"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context, userId, postId int64) models.HandleFuncResp {
	var comments []Comment.Comment
	mysql_comment.GetComments(&comments, postId)

	responseComments := make([]Comment.ResponseComment, len(comments))
	for idx, comment := range comments {
		author := mysql_user.FindAuthorInfo(comment.AuthorID)
		res := Comment.BuildResponseComment(c, userId, &comment, author)
		responseComments[idx] = res
	}

	return models.HandleFuncResp{
		Data: gin.H{
			"comments": responseComments,
			"length":   len(responseComments),
		},
		Ok: true,
	}
}
