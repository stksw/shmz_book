package main

import (
	"net/http"
	"shmz_book/handler"
	"shmz_book/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// どのハンドラーの実装をどのURLで公開するかルーティングする
func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	validator := validator.New()
	addTask := &handler.AddTask{Store: store.Tasks, Validator: validator}
	mux.Post("/tasks", addTask.ServeHTTP)
	taskList := &handler.TaskList{Store: store.Tasks}
	mux.Get("/tasks", taskList.ServeHTTP)
	return mux
}
