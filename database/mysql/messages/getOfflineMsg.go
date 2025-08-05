package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
)

func GetOfflineMsg(to int64, msgs *[]ws_model.MessageWs) {
	global.DB.Where("`state` = ? AND `to` = ?", 1, to).Find(msgs)
}

func FinishOfflineMsg(msgs *[]ws_model.MessageWs) {

	ids := make([]int64, 0, len(*msgs))
	for _, msg := range *msgs {
		ids = append(ids, msg.Id)
	}

	// 批量更新：将这些 ID 对应的消息 State 设为 false
	global.DB.Model(&ws_model.MessageWs{}).Where("id IN ?", ids).Update("state", false)
}
