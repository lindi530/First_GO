package models

import "time"

type Post struct {
	PostID    int64     `gorm:"primaryKey;column:post_id"  json:"post_id"`
	UserID    int64     `gorm:"index;not null"             json:"user_id"`
	Title     string    `gorm:"size:200;not null"          json:"title"`
	Content   string    `gorm:"type:text;not null"         json:"content"`
	Status    int8      `gorm:"default:0"                  json:"status"` // 0=正常，1=删除
	CreatedAt time.Time `gorm:"autoCreateTime"             json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"             json:"updated_at"`

	Author User `gorm:"foreignKey:UserID;references:UserID" json:"-"`
}

type AuthorInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
type PostResponse struct {
	PostID    int64      `json:"post_id"`
	UserID    int64      `json:"user_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Status    int8       `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Author    AuthorInfo `json:"author"`
}
