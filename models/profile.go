package models

type UserProfile struct {
	UserName string `json:"user_name" binding:"required"`
	Quote    string `json:"quote" binding:"max=50"`
	Email    string `json:"email"  binding:"email"`
}
