package messages

import (
	"GO1/global"
	"GO1/models/ws"
)

func GetOfflineMsg(to int64, msgs *[]ws.Message) {
	global.DB.Where("`state` = ? AND `to` = ?", 1, to).Find(msgs)
}

func ModifyOfflineMsg(msgs []ws.Message) {
	for _, msg := range msgs {
		msg.State = false
	}
	global.DB.Save(msgs)
}
