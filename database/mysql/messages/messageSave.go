package messages

import (
	"GO1/global"
	"GO1/models/ws"
	"time"
)

func MessageSave(message *ws.Message) {
	if message.Id != 0 {
		return
	}
	
	message.SendTime = time.Now()
	global.DB.Create(&message)
}
