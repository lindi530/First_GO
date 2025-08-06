package ws_model

import (
	"GO1/global"
	"encoding/json"
	"go.uber.org/zap"
)

func DecodeWs(raw []byte) *MessageWs {
	var msg MessageWs
	if err := json.Unmarshal(raw, &msg); err != nil {
		global.Logger.Warn("failed to parse message", zap.Error(err))
		return nil
	}

	switch msg.Type {
	case "chat":
		// 类型断言和转化
		m, ok := msg.Data.(map[string]interface{})
		if !ok {
			return nil
		}
		var chat ChatWS
		if b, err := json.Marshal(m); err == nil {
			_ = json.Unmarshal(b, &chat)
			msg.Data = &chat
		}
	}

	return &msg
}
