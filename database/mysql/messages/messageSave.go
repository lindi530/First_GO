package messages

import (
	"GO1/global"
	"GO1/models/ws"
	"github.com/goccy/go-json"
	"time"
)

func MessageSave(message []byte) {
	var msg ws.Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return
	}
	msg.SendTime = time.Now()
	global.DB.Create(&msg)
}
