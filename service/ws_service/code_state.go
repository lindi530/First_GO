package ws_service

import (
	"GO1/database/mysql/messages"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) CodeStateWs(msg *ws_model.MessageWs) {
	if receiver, ok := h.clients[msg.To]; ok {
		data, _ := json.Marshal(msg)
		receiver.Send <- data
	} else {

	}
	messages.MessageSave(msg)
}
