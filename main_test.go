package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	errg, ctx := errgroup.WithContext(ctx)
	errg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get %+v", err)
	}

	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body %v", err)
	}

	// http serverの戻り値を検証
	want := fmt.Sprintf("hello %s", in)
	if string(got) != want {
		t.Errorf("want %q, but %q", want, got)
	}

	// run関数に終了通知を送信
	cancel()
	// run関数の戻り値を検証
	if err := errg.Wait(); err != nil {
		t.Fatal(err)
	}

}
