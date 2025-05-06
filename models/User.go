package models

type User struct {
	Id int `gorm:"primary_key" json:"UserId" binding:"required"`
}
