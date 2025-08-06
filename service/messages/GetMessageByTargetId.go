package messages

import (
	"GO1/database/mysql/messages"
	"GO1/models"
	"GO1/models/ws_model"
)

func GetMessageByTargetId(from, to int64) (resp models.HandleFuncResp) {

	data := []ws_model.Chat{}
	err := messages.GetMessageByTargetId(&data, from, to)

	msg := make([]ws_model.MessageWs, len(data), len(data))
	for idx, d := range data {
		d.DBToWs(&msg[idx])
	}

	resp.Data = msg
	resp.Ok = err == nil

	return resp
}
