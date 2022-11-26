package service

import (
	"context"
	"fmt"
	"shmz_book/entity"
	"shmz_book/store"
)

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder
}

// entityを生成してRepositoryに渡し、結果をhandlerに返す
func (at *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	task := &entity.Task{
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := at.Repo.AddTask(ctx, at.DB, task)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return task, nil
}
