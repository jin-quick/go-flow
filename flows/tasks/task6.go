/**
 * Name: task2.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/18
 * Description:
 */

package tasks

import (
	"context"
	"errors"
	"github.com/jin-quick/go-flow/pkg/logger"
	"go.uber.org/zap"
)

type Task6 struct {
	*baseTask
}

func (t *Task6) Process(ctx context.Context, data interface{}) error {
	value, ok := data.(string)
	if !ok {
		return errors.New("invalid data type")
	}
	logger.Info(ctx, "task processing",
		zap.String("TaskName", t.Name()),
		zap.String("Value", value))
	return nil
}

func (t *Task6) Name() string {
	return "Task6"
}

func (t *Task6) Verify(data interface{}) bool {
	value, ok := data.(string)
	return ok && value == "2"
}

func NewTask6() Task {
	return &Task6{&baseTask{}}
}

// End
