package Comment

import (
	models_user "GO1/models/user"
	"time"
)

type ResponseComment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	Likes     int64     `json:"likes"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at"`

	Author models_user.AuthorInfo `json:"author"`
}
