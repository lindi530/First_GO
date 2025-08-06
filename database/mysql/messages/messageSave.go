package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
)

func ChatSave(msg *ws_model.MessageWs) {
	if msg.Data.(*ws_model.ChatWS).ID != 0 {
		return
	}

	chatData := &ws_model.Chat{}
	msg.WsToDB(chatData)

	global.DB.Create(chatData)
}
