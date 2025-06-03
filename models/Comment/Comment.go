package Comment

import (
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	AuthorID  int64     `json:"author_id"`
	Content   string    `json:"content"`
	Likes     int64     `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
