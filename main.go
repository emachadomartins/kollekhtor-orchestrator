package main

import (
	"encoding/json"
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
			creationTask, err := task.Build(msg.Body)
			if err != nil {
				fmt.Printf("Error processing message: %s\n", err.Error())
			}

			taskJSON, _ := json.Marshal(creationTask)
			fmt.Printf("Received a message: %s\n", taskJSON)
		case sig := <-sigs:
			fmt.Printf("Shutting down gracefully on signal: %s\n", sig)
			return
		}
	}
}
