package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
)

func GetOfflineMsg(to int64, msgs *[]ws_model.Chat) {
	global.DB.Where("`state` = ? AND `to` = ?", 1, to).Find(msgs)
}

func FinishOfflineMsg(msgs *[]ws_model.Chat) {

	ids := make([]int64, 0, len(*msgs))
	for _, msg := range *msgs {
		global.Logger.Info("msgID: ", msg.ID)
		ids = append(ids, msg.ID)
	}

	// 批量更新：将这些 ID 对应的消息 State 设为 false
	global.DB.Model(&ws_model.Chat{}).Where("id IN ?", ids).Update("state", false)
}
