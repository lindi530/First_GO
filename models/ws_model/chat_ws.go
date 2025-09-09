package ws_model

import (
	"strconv"
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

type ChatMessage struct {
	Type     string    `json:"type"`
	ID       int64     `json:"id"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Content  string    `json:"content"`
	State    bool      `json:"state"`
	SendTime time.Time `json:"send_time"`
}

func (c ChatMessage) Convert() Chat {
	to, _ := strconv.ParseInt(c.To, 10, 64)
	from, _ := strconv.ParseInt(c.From, 10, 64)
	return Chat{
		Type:     "chat",
		ID:       c.ID,
		From:     from,
		To:       to,
		Content:  c.Content,
		State:    c.State,
		SendTime: c.SendTime,
	}
}
