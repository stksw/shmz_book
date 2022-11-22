package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"shmz_book/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	lsn, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", lsn.Addr().String())
	log.Printf("start with: %v", url)

	// http.Handlerが返ってくる
	mux, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	// http.Handlerを元にServer構造体を生成
	s := NewServer(lsn, mux)
	// 別Goroutineで起動
	return s.Run(ctx)
}
