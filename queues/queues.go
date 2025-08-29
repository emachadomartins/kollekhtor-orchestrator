package queues

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	name       string
	url        string
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Messages   <-chan amqp.Delivery
}

func (c *Queue) Close() {
	if c.Channel != nil {
		_ = c.Channel.Close()
	}
	if c.Connection != nil {
		_ = c.Connection.Close()
	}
	fmt.Printf("Closed connection to queue `%s` in `%s`\n", c.name, c.url)
}

func NewConsumer(queueURL, queueName string) (*Queue, error) {
	fmt.Printf("Start consuming queue `%s` in `%s`\n", queueName, queueURL)
	conn, err := amqp.Dial(queueURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()
		return nil, err
	}

	return &Queue{
		name:       queueName,
		url:        queueURL,
		Connection: conn,
		Channel:    ch,
		Messages:   msgs,
	}, nil
}
