package main

import (
	"context"
	"net/http"
	"shmz_book/clock"
	"shmz_book/config"
	"shmz_book/handler"
	"shmz_book/service"
	"shmz_book/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// どのハンドラーの実装をどのURLで公開するかルーティング
func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	repo := store.Repository{Clocker: clock.RealClocker{}}
	validator := validator.New()
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &repo},
		Validator: validator,
	}
	mux.Post("/tasks", at.ServeHTTP)
	tl := &handler.TaskList{
		Service: &service.TaskList{DB: db, Repo: &repo},
	}
	mux.Get("/tasks", tl.ServeHTTP)
	return mux, cleanup, nil
}
