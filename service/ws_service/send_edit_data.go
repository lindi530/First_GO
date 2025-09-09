package ws_service

import (
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendEditData(msg *ws_model.EditStatus, content string) {
	msg.Content = content
	global.Logger.Info(content)
	if receiver, ok := h.clients[msg.To]; ok {
		data, _ := json.Marshal(msg)
		receiver.Send <- data
	} else {

	}
	//messages.MessageSave(msg)
}
