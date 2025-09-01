package comment_like_model

import "time"

type CommentLike struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"index;not null"` // 点赞用户
	CommentID int64     `gorm:"index;not null"` // 被点赞的评论
	CreatedAt time.Time // 点赞时间
}

// 设置联合唯一索引，防止重复点赞
func (CommentLike) TableName() string {
	return "comment_likes"
}
