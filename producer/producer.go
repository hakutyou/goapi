package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
)

var redis = &asynq.RedisClientOpt{
	Addr: "localhost:6379",
	// Omit if no password is required
	// Password: "mypassword",
	// Use a dedicated db number for asynq.
	// By default, Redis offers 16 databases (0..15)
	DB: 2,
}

func main() {
	client := asynq.NewClient(redis)

	t1 := asynq.NewTask(
		"send_welcome_email",
		map[string]interface{}{"user_id": 42})
	t2 := asynq.NewTask(
		"send_reminder_email",
		map[string]interface{}{"user_id": 42})

	err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}

	// 24 小时后执行
	err = client.EnqueueIn(24*time.Hour, t2)
	if err != nil {
		log.Fatal(err)
	}
}
