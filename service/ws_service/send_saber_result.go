package ws_service

import (
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendSaberResult(msg *ws_model.SaberResult) {
	if receiver, ok := h.clients[msg.To]; ok {
		data, _ := json.Marshal(msg)
		receiver.Send <- data
	} else {
		global.Logger.Error("")
	}
}
