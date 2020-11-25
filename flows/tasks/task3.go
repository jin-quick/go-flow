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

// Task3 タスク３
type Task3 struct {
	*baseTask
}

// Process タスク実行する
func (t *Task3) Process(ctx context.Context, data interface{}) error {
	value, ok := data.(string)
	if !ok {
		return errors.New("invalid data type")
	}
	logger.Info(ctx, "task processing",
		zap.String("TaskName", t.Name()),
		zap.String("Value", value))
	return nil
}

// Name タスク名
func (t *Task3) Name() string {
	return "Task3"
}

// Verify タスク実行可能かチェック
func (t *Task3) Verify(data interface{}) bool {
	value, ok := data.(string)
	return ok && value == "3"
}

// NewTask3 初期化
func NewTask3() Task {
	return &Task3{&baseTask{}}
}

// End
