package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID   int64
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	clients    map[int64]*Client
	register   chan *Client
	unregister chan *Client
	privateMsg chan *Message
	mu         sync.RWMutex
}

type Message struct {
	Type     string `json:"type"`      //
	From     int64  `json:"from"`      // 发送者ID
	To       int64  `json:"to"`        // 接收者ID
	Content  string `json:"content"`   // 内容
	ChatID   string `json:"chat_id"`   // 会话唯一ID（如: user_1_2）
	SendTime int64  `json:"send_time"` // 时间戳（客户端可选）
}
