package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/emachadomartins/kollekthor-orkhestrathor/env"
	"github.com/emachadomartins/kollekthor-orkhestrathor/queues"
	"github.com/emachadomartins/kollekthor-orkhestrathor/task"
)

func main() {
	consumer, err := queues.NewConsumer(env.QueueURL, env.QueueName)
	if err != nil {
		panic(
			errors.New(
				fmt.Sprintf(
					"error creating consumer: %s",
					err.Error(),
				),
			),
		)
	}

	defer consumer.Close()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case msg := <-consumer.Messages:
			task.Build(msg.Body)
			fmt.Printf("Received a message: %s\n", msg.Body)
		case sig := <-sigs:
			fmt.Printf("Shutting down gracefully on signal: %s\n", sig)
			return
		}
	}
}
