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
	"strconv"

	"github.com/jin-quick/go-flow/pkg/logger"
	"go.uber.org/zap"
)

// Task2 タスク２
type Task2 struct {
	*baseTask
}

// Process タスク実行
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

// Name task name
func (t *Task2) Name() string {
	return "Task2"
}

// Verify タスク実行可能かチェック
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

// NewTask2 初期化
func NewTask2() Task {
	return &Task2{&baseTask{}}
}

// End
