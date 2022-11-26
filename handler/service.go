package handler

import (
	"context"
	"shmz_book/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskListService AddTaskService
type TaskListService interface {
	TaskList(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
