package Comment

import (
	models_user "GO1/models/user"
	"github.com/gin-gonic/gin"
	"time"
)

type ResponseComment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	Author models_user.AuthorInfo `json:"author"`
}

func BuildResponseComment(c *gin.Context, comment *Comment, user *models_user.AuthorInfo) ResponseComment {
	return ResponseComment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Author: models_user.AuthorInfo{
			UserID:   user.UserID,
			UserName: user.UserName,
			Avatar:   models_user.GetAvatarPath(c, user.Avatar),
		},
	}
}
