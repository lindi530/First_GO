package Comment

import (
	"GO1/database/mysql/comment_like"
	"GO1/models/Comment"
	"GO1/models/user_model"
	"github.com/gin-gonic/gin"
)

func BuildResponseComment(c *gin.Context, userId int64, comment *Comment.Comment, author *user_model.AuthorInfo) Comment.ResponseComment {
	return Comment.ResponseComment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Likes:     comment.Likes,
		Like:      comment_like.CommentLikeCheck(userId, comment.ID),
		Author: user_model.AuthorInfo{
			UserID:   author.UserID,
			UserName: author.UserName,
			Avatar:   user_model.GetAvatarPath(c, author.Avatar),
		},
	}
}
