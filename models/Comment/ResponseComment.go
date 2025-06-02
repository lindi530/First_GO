package Comment

import (
	"GO1/database/mysql/comment_like"
	models_user "GO1/models/user"
	"github.com/gin-gonic/gin"
	"time"
)

type ResponseComment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at"`

	Author models_user.AuthorInfo `json:"author"`
}

func BuildResponseComment(c *gin.Context, userId int64, comment *Comment, author *models_user.AuthorInfo) ResponseComment {
	return ResponseComment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Like:      comment_like.CommentLikeCheck(userId, comment.ID),
		Author: models_user.AuthorInfo{
			UserID:   author.UserID,
			UserName: author.UserName,
			Avatar:   models_user.GetAvatarPath(c, author.Avatar),
		},
	}
}
