package handler

import (
	"net/http"
	"shmz_book/entity"
	"shmz_book/store"
)

type TaskList struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (tl *TaskList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := tl.Store.All()
	rsp := []task{}

	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	ResponseJSON(ctx, w, rsp, http.StatusOK)
}
