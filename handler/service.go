package handler

import (
	"context"
	"shmz_book/entity"
)

type TaskListService interface {
	TaskList(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
