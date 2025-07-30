package ws_service

import (
	"GO1/models/ws_model"
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
	privateMsg chan *ws_model.MessageWs
	mu         sync.RWMutex
}
