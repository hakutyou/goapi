package main

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

type Dispatcher struct {
	mapping map[string]func(*asynq.Task) error
}

func (d *Dispatcher) HandleFunc(taskType string, fn func(*asynq.Task) error) {
	d.mapping[taskType] = fn
}

func (d *Dispatcher) ProcessTask(_ context.Context, task *asynq.Task) error {
	fn, ok := d.mapping[task.Type]
	if !ok {
		return fmt.Errorf("no handler registered for %q", task.Type)
	}
	return fn(task)
}

func main() {
	if err := LoadConfigure(); err != nil {
		panic(err)
	}
	if err := openRedis(); err != nil {
		panic(err)
	}

	d := &Dispatcher{mapping: make(map[string]func(*asynq.Task) error)}
	tasks(d)

	bg := asynq.NewServer(redis, asynq.Config{
		Concurrency: 10,
	})
	_ = bg.Run(d)
}
