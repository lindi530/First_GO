package messages

import (
	"GO1/global"
	"GO1/models/ws_model"
)

func GetMessageByTargetId(msg *[]ws_model.Chat, from, to int64) error {
	err := global.DB.
		Where("(`from` = ? AND `to` = ?) OR (`from` = ? AND `to` = ?)", from, to, to, from).
		Order("send_time ASC").
		Find(msg).Error

	return err
}
