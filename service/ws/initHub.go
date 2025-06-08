package ws

import "GO1/models/ws"

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		privateMsg: make(chan *ws.Message),
	}
}

var WsHub = NewHub()
