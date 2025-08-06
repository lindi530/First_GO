package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendData(msg *ws_model.MessageWs) {

	if receiver, ok := h.clients[msg.To]; ok {
		// 你可以用 json.Marshal 来编码成 []byte
		data, _ := json.Marshal(msg.Data)
		receiver.Send <- data
	} else {
		if data, ok := msg.Data.(*ws_model.ChatWS); ok {
			data.State = true
		}
	}
	global.Logger.Info(msg.Data.(*ws_model.ChatWS).State)
	messages.ChatSave(msg)
}
