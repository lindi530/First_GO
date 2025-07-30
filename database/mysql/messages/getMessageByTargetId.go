package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
)

func GetMessageByTargetId(from, to int64) ([]ws_model.MessageWs, error) {
	var msg []ws_model.MessageWs
	err := global.DB.
		Where("(`from` = ? AND `to` = ?) OR (`from` = ? AND `to` = ?)", from, to, to, from).
		Order("send_time ASC").
		Find(&msg).Error

	return msg, err
}
