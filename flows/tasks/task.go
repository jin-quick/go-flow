/**
 * Name: task.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/18
 * Description:
 */

package tasks

import "context"

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

func (b *baseTask) Add(task Task) {
	b.tasks = append(b.tasks, task)
}

func (b *baseTask) GetNext(data interface{}) Task {
	for _, task := range b.tasks {
		if task.Verify(data) {
			return task
		}
	}
	return nil
}

// End
