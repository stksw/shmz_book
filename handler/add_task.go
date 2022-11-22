package handler

import (
	"encoding/json"
	"net/http"
	"shmz_book/entity"

	"github.com/go-playground/validator/v10"
)

type AddTask struct {
	Service   AddTaskService
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		Title string `json:"title" validate:"required"`
	}

	// JSONをデコードしてbodyにセット
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// バリデーション
	err := at.Validator.Struct(body)
	if err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
	}

	// タスクを追加
	task, err := at.Service.AddTask(ctx, body.Title)
	if err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: task.ID}
	ResponseJSON(ctx, w, rsp, http.StatusOK)
}
