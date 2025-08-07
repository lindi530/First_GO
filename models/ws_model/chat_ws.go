package ws_model

import (
	"time"
)

type Chat struct {
	Type     string    `json:"type" gorm:"-"`
	ID       int64     `json:"id"`
	From     int64     `json:"from"`
	To       int64     `json:"to"`
	Content  string    `json:"content"`
	State    bool      `json:"state"`
	SendTime time.Time `json:"send_time"`
}
