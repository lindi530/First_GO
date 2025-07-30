package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
	"time"
)

func MessageSave(message *ws_model.MessageWs) {
	if message.Id != 0 {
		return
	}

	message.SendTime = time.Now()
	global.DB.Create(message)
}
