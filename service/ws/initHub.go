package ws

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int64]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		privateMsg: make(chan *Message),
	}
}

var WsHub = NewHub()
