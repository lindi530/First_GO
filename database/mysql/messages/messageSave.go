package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
	"time"
)

func ChatSave(chatData *ws_model.Chat) {
	if chatData.ID != 0 {
		return
	}
	chatData.SendTime = time.Now()
	global.DB.Create(chatData)
}
