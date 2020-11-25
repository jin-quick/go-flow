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
	"strconv"
)

type Task2 struct {
	*baseTask
}

func (t *Task2) Process(ctx context.Context, data interface{}) error {
	value, ok := data.(string)
	if !ok {
		return errors.New("invalid data type")
	}
	logger.Info(ctx, "task processing",
		zap.String("TaskName", t.Name()),
		zap.String("Value", value))
	return nil
}

func (t *Task2) Name() string {
	return "Task2"
}

func (t *Task2) Verify(data interface{}) bool {
	value, ok := data.(string)
	if ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return i%2 == 0
	}
	return false
}

func NewTask2() Task {
	return &Task2{&baseTask{}}
}

// End
