package ws_service

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
	privateMsg chan []byte
	mu         sync.RWMutex
}
