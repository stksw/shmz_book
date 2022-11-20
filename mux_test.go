package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// ResponseWriterインターフェースを満たす *ResponseRecorder型の値を取得
	w := httptest.NewRecorder()
	// テスト用の*http.Requestを生成
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := NewMux()

	// ServeHTTPに渡した後、Resultを実行するとクライアントが受け取る
	// レスポンスが含まれるhttp.Response型を取得
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read body %v", err)
	}

	want := `{"status": " ok"}`
	if string(got) != want {
		t.Errorf("want %q, but %q", want, got)
	}

}
