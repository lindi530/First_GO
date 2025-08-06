package ws_service

import (
	"GO1/models/ws_model"
)

func (h *Hub) Run() {
	for {
		select {
		case msg := <-h.register:
			h.mu.Lock()
			h.clients[msg.ID] = msg
			h.mu.Unlock()
		case msg := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[msg.ID]; ok {
				delete(h.clients, msg.ID)
				close(msg.Send)
			}
			h.mu.Unlock()
		case msg := <-h.privateMsg:
			h.mu.RLock()
			switch msg.Type {
			case "chat":
				h.SendData(msg)
				break
			case "submit_code":
				break
			default:
				break
			}

			h.mu.RUnlock()
		}
	}
}

// 导出封装方法
func (h *Hub) RegisterClient(client *Client) {
	h.register <- client
}

func (h *Hub) UnregisterClient(client *Client) {
	h.unregister <- client
}

func (h *Hub) Broadcast(message *ws_model.MessageWs) { // 广播
	h.privateMsg <- message
}
