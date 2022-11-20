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

	mux := NewMux()
	s := NewServer(lsn, mux)
	return s.Run(ctx)
}
