package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	lsn net.Listener
}

func NewServer(lsn net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		lsn: lsn,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// 終了シグナルを待機
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	errg, ctx := errgroup.WithContext(ctx)
	// 別Goroutineでhttp serverを起動
	errg.Go(func() error {
		if err := s.srv.Serve(s.lsn); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	// チャネルからの終了通知を待機
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %+v", err)
	}

	// errg.Goで起動したGoroutineの終了を待つ
	return errg.Wait()
}
