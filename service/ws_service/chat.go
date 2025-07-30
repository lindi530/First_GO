package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) ChatWs(msg *ws_model.MessageWs) {
	if receiver, ok := h.clients[msg.To]; ok {
		// 你可以用 json.Marshal 来编码成 []byte
		data, _ := json.Marshal(msg)
		receiver.Send <- data
	} else {
		msg.State = true
	}
	messages.MessageSave(msg)
}
