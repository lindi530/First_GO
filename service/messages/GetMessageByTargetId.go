package messages

import (
	"GO1/database/mysql/messages"
	"GO1/models"
)

func GetMessageByTargetId(from, to int64) (resp models.HandleFuncResp) {
	data, err := messages.GetMessageByTargetId(from, to)
	resp.Data = data
	resp.Ok = err == nil
	
	return resp
}
