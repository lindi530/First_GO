package models

type UserProfile struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email"  binding:"email"`
}
