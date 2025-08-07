package messages

import (
	"GO1/database/mysql/messages"
	"GO1/models"
	"GO1/models/ws_model"
)

func GetMessageByTargetId(from, to int64) (resp models.HandleFuncResp) {

	msg := []ws_model.Chat{}
	err := messages.GetMessageByTargetId(&msg, from, to)

	resp.Data = &msg
	resp.Ok = err == nil

	return resp
}
