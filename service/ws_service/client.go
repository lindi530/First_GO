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
		_, typeDataWs, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		hub.Broadcast(typeDataWs)
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
	var msgs []ws_model.Chat
	messages.GetOfflineMsg(userId, &msgs)

	for _, msg := range msgs {
		msg.Type = ws_model.MessageTypeChat
		data, err := json.Marshal(msg)
		if err != nil {
			continue
		}
		hub.Broadcast(data)
	}
	messages.FinishOfflineMsg(&msgs)
}
