package mq

import (
	"fmt"
	"todo_list/config"

	"github.com/streadway/amqp"
)

var RabbitMq *amqp.Connection

func InitRabbitMQ() {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/",
		config.RabbitMQ,
		config.RabbitMQUser,
		config.RabbitMQPassword,
		config.RabbitMQHost,
		config.RabbitMQPort)
	fmt.Println(connString)
	conn, err := amqp.Dial(connString)
	if err != nil {
		RabbitMq = conn
	}
}
