package post

import (
	"GO1/models/user"
	"github.com/gin-gonic/gin"
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
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
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

func BuildPostResponse(c *gin.Context, p Post) PostResponse {
	return PostResponse{
		PostID:    p.PostID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Author: AuthorInfo{
			UserId:   p.Author.UserID,
			UserName: p.Author.UserName,
			Avatar:   user.GetAvatarPath(c, p.Author.Avatar),
		},
	}
}

func BuildPostsResponse(c *gin.Context, p []Post) []PostResponse {
	postsResponse := make([]PostResponse, 0, len(p))
	for _, post := range p {
		postsResponse = append(postsResponse, BuildPostResponse(c, post))
	}
	return postsResponse
}
