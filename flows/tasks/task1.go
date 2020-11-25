/**
 * Name: task1.go
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

type Task1 struct {
	*baseTask
}

func (t *Task1) Process(ctx context.Context, data interface{}) error {
	value, ok := data.(string)
	if !ok {
		return errors.New("invalid data type")
	}
	logger.Info(ctx, "task processing",
		zap.String("TaskName", t.Name()),
		zap.String("Value", value))
	return nil
}

func (t *Task1) Name() string {
	return "Task1"
}

func (t *Task1) Verify(_ interface{}) bool {
	return true
}

func NewTask1() Task {
	return &Task1{&baseTask{}}
}

// End
