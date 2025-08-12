package Comment

import (
	"GO1/models/user_model"
	"time"
)

type ResponseComment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	Content   string    `json:"content"`
	Likes     int64     `json:"likes"`
	Like      bool      `json:"like"`
	CreatedAt time.Time `json:"created_at"`

	Author user_model.AuthorInfo `json:"author"`
}
