package service

import (
	"context"
	"fmt"
	"shmz_book/entity"
	"shmz_book/store"
)

type TaskList struct {
	DB   store.Queryer
	Repo TaskLister
}

// Repositoryに渡して、結果をhandlerに返す
func (tl *TaskList) TaskList(ctx context.Context) (entity.Tasks, error) {
	ts, err := tl.Repo.TaskList(ctx, tl.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
