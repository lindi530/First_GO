package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func (c *Client) ReadLoop(hub *Hub) {
	defer func() {
		hub.UnregisterClient(c)
		if err := c.Conn.Close(); err != nil {
			global.Logger.Error("Close error: %v", err)
		}
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		var msg ws_model.MessageWs

		if err := json.Unmarshal(message, &msg); err != nil {
			continue // 无效消息
		}
		msg.State = false
		global.Logger.Warn("online", msg)
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

func (c *Client) WriteOfflineMsg(hub *Hub, userId int64) {
	var msgs = []ws_model.MessageWs{}
	messages.GetOfflineMsg(userId, &msgs)

	for _, msg := range msgs {
		hub.Broadcast(&msg)
	}
	messages.FinishOfflineMsg(&msgs)
}
