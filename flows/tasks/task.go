/**
 * Name: task.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/18
 * Description:
 */

package tasks

import "context"

// Task タスク
type Task interface {
	Process(ctx context.Context, data interface{}) error
	Name() string
	Add(task Task)
	Verify(data interface{}) bool
	GetNext(data interface{}) Task
}

type baseTask struct {
	tasks []Task
}

// Add サブタスク追加
func (b *baseTask) Add(task Task) {
	b.tasks = append(b.tasks, task)
}

// GetNext 次のタスクを取得
func (b *baseTask) GetNext(data interface{}) Task {
	for _, task := range b.tasks {
		if task.Verify(data) {
			return task
		}
	}
	return nil
}

// End
