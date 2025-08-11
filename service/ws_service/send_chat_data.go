package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendChatData(data []byte) {

	var chatData ws_model.Chat
	if err := json.Unmarshal(data, &chatData); err != nil {
		return
	}

	if receiver, ok := h.clients[chatData.To]; ok {
		receiver.Send <- data
	} else {
		chatData.State = true
	}
	//global.Logger.Info(msg.Data.(*ws_model.ChatWS).State)

	messages.ChatSave(&chatData)
}
