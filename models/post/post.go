package post

import (
	"GO1/models/user"
	"time"
)

// GORM 外键绑定 + Preload     评论是用手动查询来写(更推荐)
type Post struct {
	PostID    int64     `gorm:"primaryKey;column:id"  json:"id,omitempty"`
	UserID    int64     `gorm:"index;not null"             json:"user_id"`
	Title     string    `gorm:"size:200;not null"          json:"title"`
	Content   string    `gorm:"type:text;not null"         json:"content"`
	Status    int8      `gorm:"default:0"                  json:"status"` // 0=正常，1=删除
	CreatedAt time.Time `gorm:"autoCreateTime"             json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"             json:"updated_at"`

	Author user.User `gorm:"foreignKey:UserID;references:UserID" json:"-"`
}

type AuthorInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
type PostResponse struct {
	PostID    int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
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

func BuildPostResponse(p Post) PostResponse {
	return PostResponse{
		PostID:    p.PostID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Author: AuthorInfo{
			UserName: p.Author.UserName,
			Email:    p.Author.Email,
		},
	}
}

func BuildPostsResponse(p []Post) []PostResponse {
	postsResponse := make([]PostResponse, 0, len(p))
	for _, post := range p {
		postsResponse = append(postsResponse, BuildPostResponse(post))
	}
	return postsResponse
}
