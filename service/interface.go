package service

import (
	"context"
	"shmz_book/entity"
	"shmz_book/store"
)

// service.AddTaskのRepoに適用する
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

// service.TaskListのRepoに適用する
type TaskLister interface {
	TaskList(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
