package Comment

import (
	"GO1/database/mysql/comment_like"
	"GO1/models/Comment"
	models_user "GO1/models/user"
	"github.com/gin-gonic/gin"
)

func BuildResponseComment(c *gin.Context, userId int64, comment *Comment.Comment, author *models_user.AuthorInfo) Comment.ResponseComment {
	return Comment.ResponseComment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Likes:     comment.Likes,
		Like:      comment_like.CommentLikeCheck(userId, comment.ID),
		Author: models_user.AuthorInfo{
			UserID:   author.UserID,
			UserName: author.UserName,
			Avatar:   models_user.GetAvatarPath(c, author.Avatar),
		},
	}
}
