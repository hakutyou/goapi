package main

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func sendWelcomeEmail(t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to User %d\n", id)
	return nil
}

func sendReminderEmail(t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to User %d\n", id)
	return nil
}
