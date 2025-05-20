package models

import (
	"time"
)

type User struct {
	UserID     int64     `json:"user_id"`
	UserName   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type UserResponse struct {
	UserID     int64     `json:"user_id"`
	UserName   string    `json:"username"`
	Email      string    `json:"email"`
	Gender     string    `json:"gender"`
	UpdateTime time.Time `json:"update_time"`
}

func BuildUserResponse(u User) UserResponse {
	user := UserResponse{
		UserID:     u.UserID,
		UserName:   u.UserName,
		Email:      u.Email,
		Gender:     u.Gender,
		UpdateTime: u.UpdateTime,
	}
	return user
}
