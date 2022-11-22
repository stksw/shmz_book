package handler

import (
	"net/http"
	"shmz_book/entity"
)

type TaskList struct {
	Service TaskListService
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (tl *TaskList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := tl.Service.TaskList(ctx)
	// tasks, err := tl.Repo.TaskList(ctx, tl.DB)
	if err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

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
