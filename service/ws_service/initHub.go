package ws_service

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		privateMsg: make(chan []byte),
	}
}

var WsHub = NewHub()
