package store

import (
	"errors"
	"fmt"
	"shmz_book/entity"
)

var (
	Tasks       = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}
	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	fmt.Println("len", len(ts.Tasks))
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i] = t
	}
	return tasks
}
