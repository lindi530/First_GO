package post

import (
	database_comment "GO1/database/mysql/comment"
	"GO1/models"
	"GO1/models/Comment"
	"fmt"
	"time"
)

func CreateComment(userId, postId int64, comment *Comment.Comment) models.HandleFuncResp {
	comment.PostID = postId
	comment.AuthorID = userId
	comment.CreatedAt = time.Now()

	fmt.Println(comment)
	result := database_comment.CreateComment(comment)

	return result
}
