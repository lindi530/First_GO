package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendChatData(data []byte) {

	var chatMessage ws_model.ChatMessage
	if err := json.Unmarshal(data, &chatMessage); err != nil {
		global.Logger.Error(err)
		return
	}

	chatData := chatMessage.Convert()
	global.Logger.Info(chatData)

	if receiver, ok := h.clients[chatData.To]; ok {
		receiver.Send <- data
	} else {
		chatData.State = true
	}
	//global.Logger.Info(msg.Data.(*ws_model.ChatWS).State)

	messages.ChatSave(&chatData)
}
