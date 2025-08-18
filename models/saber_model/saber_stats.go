package saber_model

import "time"

type SaberStat struct {
	UserID       int64     `json:"user_id"`
	Rating       int       `json:"rating"`
	Level        string    `json:"level"`
	Wins         int       `json:"wins"`
	TotalMatches int       `json:"total_matches"`
	CreatedAT    time.Time `json:"created_at"`
	UpdatedAT    time.Time `json:"updated_at"`
}
