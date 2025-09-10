package rabbitmq

import (
	"GO1/global"
	"github.com/streadway/amqp"
	"log"
	"os"
)

var MQConn *amqp.Connection
var MQChannel *amqp.Channel

func InitRabbitMQ() {

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = global.Config.RabbitMQ.URL
	}

	var err error
	MQConn, err = amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("RabbitMQ connect error: %v", err)
	}

	MQChannel, err = MQConn.Channel()
	if err != nil {
		log.Fatalf("RabbitMQ channel error: %v", err)
	}

	_, err = MQChannel.QueueDeclare(
		"match_queue", // 队列名
		true,          // 队列持久化，服务器重启后仍存在
		false,         // 是否在无人使用时自动删除，false 表示不删除
		false,         // 是否排他队列，只能被当前连接使用，false 表示否
		false,         // 是否异步，false 表示同步等待服务器确认
		nil,           // 额外参数，这里不传
	)
	if err != nil {
		global.Logger.Error("RabbitMQ queue declare error: %v", err)
	}
	global.MQConn = MQConn
	global.MQChannel = MQChannel

	global.Logger.Info("RabbitMQ initialized.")
}
