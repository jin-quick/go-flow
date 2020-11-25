/**
 * Name: flow.go
 * Created by: Mengzhuang Jin
 * Created on: 2020/11/18
 * Description:
 */

package flows

import (
	"context"

	"github.com/jin-quick/go-flow/flows/tasks"
	"github.com/jin-quick/go-flow/pkg/logger"
	"go.uber.org/zap"
)

// Flow フロー
type Flow struct {
	root tasks.Task

	tasks []tasks.Task
}

// Add タスク追加
func (f *Flow) Add(parentTask, task tasks.Task) {
	f.tasks = append(f.tasks, task)
	parentTask.Add(task)
}

// Execute フローを実行する
func (f *Flow) Execute(ctx context.Context, data string) error {
	task := f.root

	if data == "4" {
		// execute all tasks
		if err := f.root.Process(ctx, data); err != nil {
			return err
		}
		for _, t := range f.tasks {
			if t.Verify(data) {
				if err := t.Process(ctx, data); err != nil {
					return err
				}
			}
		}
		return nil
	}

	for task != nil && task.Verify(data) {
		logger.Info(ctx, "task process", zap.String("TaskName", task.Name()))
		if err := task.Process(ctx, data); err != nil {
			return err
		}
		task = task.GetNext(data)
	}
	return nil
}

// NewInquirySupportFlow 初期化
func NewInquirySupportFlow() *Flow {
	f := &Flow{}

	root := tasks.NewTask1()
	f.root = root

	task2 := tasks.NewTask2()
	task3 := tasks.NewTask3()
	task4 := tasks.NewTask4()

	f.Add(root, task2)
	f.Add(root, task3)
	f.Add(root, task4)

	f.Add(task2, tasks.NewTask5())
	f.Add(task4, tasks.NewTask6())

	return f
}

// End
