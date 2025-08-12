package match_service

import (
	"GO1/database/mysql/match_mysql"
	"GO1/global"
	"GO1/models/match_model"
	"encoding/json"
	"github.com/streadway/amqp"
)

func SendMatchRequest(userid int64) error {
	user := match_model.MatchUser{
		UserID: userid,
	}
	match_mysql.GetMatchUser(&user)

	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return global.MQChannel.Publish(
		"",            // exchange
		"match_queue", // routing key (queue name)
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
}
