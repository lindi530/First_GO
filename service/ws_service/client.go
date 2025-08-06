package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/global"
	"GO1/models/ws_model"
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
		_, messageWs, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		msg := ws_model.DecodeWs(messageWs)
		if msg == nil {
			continue // 无效消息
		}

		if data, ok := msg.Data.(*ws_model.ChatWS); ok {
			data.State = false
		}

		hub.Broadcast(msg)
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
		msgWs := &ws_model.MessageWs{}
		msg.DBToWs(msgWs)
		hub.Broadcast(msgWs)
	}
	messages.FinishOfflineMsg(&msgs)
}
