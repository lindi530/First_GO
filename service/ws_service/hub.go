package ws_service

import (
	"GO1/global"
	"github.com/buger/jsonparser"
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
			h.unregisterHandle(msg)
			h.mu.Unlock()
		case msg := <-h.privateMsg:
			h.mu.RLock()

			h.sendMessage(msg)

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

func (h *Hub) Broadcast(message []byte) { // 广播
	h.privateMsg <- message
}

func (h *Hub) unregisterHandle(msg *Client) {
	if _, ok := h.clients[msg.ID]; ok {
		delete(h.clients, msg.ID)
		close(msg.Send)
	}

	h.SendOnlineData(msg.ID, false)
}

func (h *Hub) sendMessage(msg []byte) {
	msgType, err := jsonparser.GetString(msg, "type")
	if err != nil {
		global.Logger.Error("数据提取失败")
		return
	}

	switch msgType {
	case "chat":
		h.SendChatData(msg)
		break
	case "submit_code":
		break
	default:
		break
	}
}
