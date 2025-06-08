package messages

import (
	"GO1/global"
	"GO1/models/ws"
)

func GetMessageByTargetId(from, to int64) ([]ws.Message, error) {
	var msg []ws.Message
	err := global.DB.
		Where("(`from` = ? AND `to` = ?) OR (`from` = ? AND `to` = ?)", from, to, to, from).
		Order("send_time ASC").
		Find(&msg).Error

	return msg, err
}
