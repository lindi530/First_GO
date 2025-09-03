package post_model

import "time"

type PostLike struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	PostId    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
