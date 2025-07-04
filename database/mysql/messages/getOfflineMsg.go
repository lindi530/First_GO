package messages

import (
	"GO1/global"
	"GO1/models/ws"
)

func GetOfflineMsg(to int64, msgs *[]ws.Message) {
	global.DB.Where("`state` = ? AND `to` = ?", 1, to).Find(msgs)
}

func FinishOfflineMsg(msgs *[]ws.Message) {
	global.Logger.Info("ModifyOfflineMsg: ", msgs)

	ids := make([]int64, 0, len(*msgs))
	for _, msg := range *msgs {
		ids = append(ids, msg.Id)
	}

	// 批量更新：将这些 ID 对应的消息 State 设为 false
	global.DB.Model(&ws.Message{}).Where("id IN ?", ids).Update("state", false)
}
