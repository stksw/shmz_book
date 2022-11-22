package service

import (
	"context"
	"shmz_book/entity"
	"shmz_book/store"
)

type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	TaskList(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
