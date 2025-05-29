package user

import (
	"GO1/database/redis"
	"GO1/global"
	"GO1/mapping"
	models_upload "GO1/models/upload"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type User struct {
	UserID         int64     `json:"user_id" gorm:"column:user_id"`
	UserName       string    `json:"user_name" gorm:"column:user_name"`
	Password       string    `json:"password" gorm:"column:password"`
	Avatar         string    `json:"avatar" gorm:"column:avatar"`
	Email          string    `json:"email" gorm:"column:email"`
	Quote          string    `json:"quote" gorm:"column:quote"`
	FollowingCount int64     `json:"following_count" gorm:"column:following_count"`
	FollowerCount  int64     `json:"follower_count" gorm:"column:follower_count"`
	Gender         string    `json:"gender" gorm:"column:gender"`
	CreateTime     time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime     time.Time `json:"update_time" gorm:"column:update_time"`
}

type UserResponse struct {
	UserID         int64     `json:"user_id"`
	UserName       string    `json:"user_name"`
	AvatarPath     string    `json:"avatar"`
	Quote          string    `json:"quote"`
	Email          string    `json:"email"`
	FollowingCount int64     `json:"following_count"`
	FollowerCount  int64     `json:"follower_count"`
	Gender         string    `json:"gender"`
	OnlineState    bool      `json:"online_state"`
	UpdateTime     time.Time `json:"update_time"`
}

func BuildUserResponse(c *gin.Context, u User) UserResponse {
	user := UserResponse{
		UserID:         u.UserID,
		UserName:       u.UserName,
		AvatarPath:     GetAvatarPath(c, u.Avatar),
		Quote:          u.Quote,
		FollowingCount: u.FollowingCount,
		FollowerCount:  u.FollowerCount,
		Email:          u.Email,
		Gender:         u.Gender,
		OnlineState:    getOnlineState(c, u.UserID),
		UpdateTime:     u.UpdateTime,
	}
	return user
}

func getOnlineState(c *gin.Context, userId int64) bool {
	return redis.GetOnlineState(c, userId)
}

func GetAvatarPath(c *gin.Context, md5Avatar string) string {

	var path string
	global.DB.Model(models_upload.Image{}).Where("md5 = ?", md5Avatar).Select("path").First(&path)
	idx := strings.LastIndex(path, "/")
	pathHead, pathTail := path[:idx], path[idx+1:]
	scheme := "http"
	host := c.Request.Host // 如 127.0.0.1:8000 或 192.168.1.10:8000
	imageURL := scheme + "://" + host + mapping.GetRealPath(pathHead) + "/" + pathTail

	return imageURL
}
