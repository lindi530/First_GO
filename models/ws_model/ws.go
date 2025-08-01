package ws_model

import "time"

type MessageWs struct {
	Id       int64     `json:"id"`
	From     int64     `json:"from"`    // 发送者ID
	To       int64     `json:"to"`      // 接收者ID
	Content  string    `json:"content"` // 内容
	Type     string    `json:"type"`    //
	State    bool      `json:"state"`
	SendTime time.Time `json:"send_time"` // 时间戳（客户端可选）
}
