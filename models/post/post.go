package post

import (
	"GO1/models/user_model"
	"time"
)

// GORM 外键绑定 + Preload     评论是用手动查询来写(更推荐)
type Post struct {
	PostID    int64     `gorm:"primaryKey;column:id"  json:"id,omitempty"`
	UserID    int64     `gorm:"index;not null"             json:"user_id"`
	Title     string    `gorm:"size:200;not null"          json:"title"`
	Content   string    `gorm:"type:text;not null"         json:"content"`
	Likes     int64     `json:"likes"`
	Views     int64     `gorm:"not null"   json:"views"`
	Status    int8      `gorm:"default:0"                  json:"status"` // 0=正常，1=删除
	CreatedAt time.Time `gorm:"autoUpdateTime:false" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false" json:"updated_at"`

	Author user_model.User `gorm:"foreignKey:UserID;references:UserID" json:"-"`
}

type AuthorInfo struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}
type PostResponse struct {
	PostID    int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Likes     int64      `json:"likes"`
	Like      bool       `json:"like"`
	Views     int64      `json:"views"`
	Status    int8       `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Author    AuthorInfo `json:"author"`
}
type CreatePost struct {
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
