package match_service

import (
	"GO1/global"
	"GO1/models/match_model"
	"encoding/json"
)

func StartMatchConsumer() { // 监听匹配信息

	msgs, err := global.MQChannel.Consume(
		"match_queue", "", true, false, false, false, nil,
	)
	if err != nil {
		global.Logger.Error("Failed to register RabbitMQ consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			var user match_model.MatchUser
			if err := json.Unmarshal(d.Body, &user); err != nil {
				global.Logger.Error("Invalid match_model message:", err)
				continue
			}

			processMatch(&user) // 进行匹配
		}
	}()
}
