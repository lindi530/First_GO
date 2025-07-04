package ws

import (
	"GO1/database/mysql/messages"
	"GO1/global"
	"GO1/models/ws"
	"encoding/json"
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
			if msg.Type == "chat" {
				if receiver, ok := h.clients[msg.To]; ok {
					// 你可以用 json.Marshal 来编码成 []byte
					data, _ := json.Marshal(msg)
					receiver.Send <- data

					messages.MessageSave(msg)
				} else {
					msg.State = true
				}
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

func (h *Hub) Broadcast(message *ws.Message) { // 广播

	global.Logger.Info(message)
	h.privateMsg <- message
}
