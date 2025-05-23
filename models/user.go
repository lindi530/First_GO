package models

import (
	"GO1/global"
	"GO1/mapping"
	models_upload "GO1/models/upload"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type User struct {
	//ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type UserResponse struct {
	UserID     int64     `json:"user_id"`
	UserName   string    `json:"username"`
	AvatarPath string    `json:"avatar"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	UpdateTime time.Time `json:"update_time"`
}

func BuildUserResponse(c *gin.Context, u User) UserResponse {
	user := UserResponse{
		UserID:     u.UserID,
		UserName:   u.UserName,
		AvatarPath: getAvatarPath(c, u.Avatar),
		Email:      u.Email,
		Gender:     u.Gender,
		UpdateTime: u.UpdateTime,
	}
	return user
}

func getAvatarPath(c *gin.Context, md5Avatar string) string {

	var path string
	global.DB.Model(models_upload.Image{}).Where("md5 = ?", md5Avatar).Select("path").First(&path)
	idx := strings.LastIndex(path, "/")
	pathHead, pathTail := path[:idx], path[idx+1:]
	scheme := "http"
	host := c.Request.Host // 如 127.0.0.1:8000 或 192.168.1.10:8000
	imageURL := scheme + "://" + host + mapping.GetRealPath(pathHead) + "/" + pathTail

	return imageURL
}
