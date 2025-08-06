package ws_model

import (
	"time"
)

type ChatWS struct {
	ID       int64     `json:"id"`
	From     int64     `json:"from"`
	Content  string    `json:"content"`
	SendTime time.Time `json:"send_time"` // 时间戳（客户端可选）
	State    bool      `json:"state"`
}

type Chat struct {
	ID       int64     `json:"id"`
	From     int64     `json:"from"`
	To       int64     `json:"to"`
	Content  string    `json:"content"`
	State    bool      `json:"state"`
	SendTime time.Time `json:"send_time"`
}

func (msg *MessageWs) WsToDB(c *Chat) {

	chatWs := msg.Data.(*ChatWS)

	c.To = msg.To
	c.From = chatWs.From
	c.Content = chatWs.Content
	c.State = chatWs.State
	c.SendTime = time.Now()
}

func (msg *Chat) DBToWs(msgWs *MessageWs) {
	msgWs.To = msg.To
	msgWs.Type = MessageTypeChat
	msgWs.Data = &ChatWS{
		ID:       msg.ID,
		From:     msg.From,
		Content:  msg.Content,
		SendTime: msg.SendTime,
		State:    msg.State,
	}
}
