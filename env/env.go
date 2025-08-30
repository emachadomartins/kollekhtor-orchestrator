package env

import (
	"errors"
	dotenv "github.com/joho/godotenv"
	"os"
)

type environment struct {
	QueueURL  string
	QueueName string
}

func loadEnv() environment {
	_ = dotenv.Load()
	queueURL := os.Getenv("QUEUE_URL")

	if queueURL == "" {
		panic(
			errors.New("QUEUE_URL env variable not set"),
		)
	}

	queueName := os.Getenv("QUEUE_NAME")

	if queueName == "" {
		queueName = "orkhestrathor"
	}

	return environment{
		QueueURL:  queueURL,
		QueueName: queueName,
	}
}

var Env = loadEnv()

var QueueURL = Env.QueueURL

var QueueName = Env.QueueName
