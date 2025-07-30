package ws_service

import "GO1/models/ws_model"

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		privateMsg: make(chan *ws_model.MessageWs),
	}
}

var WsHub = NewHub()
