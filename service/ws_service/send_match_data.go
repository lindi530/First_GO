package ws_service

import (
	"GO1/global"
	"GO1/models/ws_model"
	"encoding/json"
)

func (h *Hub) SendMatchData(msg *ws_model.MatchResponse) {
	if receiver, ok := h.clients[msg.To]; ok {
		data, _ := json.Marshal(msg)
		receiver.Send <- data
	} else {
		global.Logger.Error("消息解析失败")
	}
}
