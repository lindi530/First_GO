package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

func (c *Client) ReadLoop(hub *Hub) {
	defer func() {
		hub.UnregisterClient(c)
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		var msg Message

		if err := json.Unmarshal(message, &msg); err != nil {
			continue // 无效消息
		}

		hub.Broadcast(&msg)
	}
}

func (c *Client) WriteLoop() {
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
