package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"shmz_book/entity"
	"shmz_book/testutil"
	"testing"
)

func TestTaskList(t *testing.T) {
	t.Parallel()

	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		tasks []*entity.Task
		want  want
	}{
		"ok": {
			tasks: []*entity.Task{
				{
					ID:     1,
					Title:  "test1",
					Status: entity.TaskStatusTodo,
				},
				{
					ID:     2,
					Title:  "test2",
					Status: entity.TaskStatusDone,
				},
			},
		},
		"empty": {
			tasks: []*entity.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			moq := &TaskListServiceMock{}
			moq.TaskListFunc = func(ctx context.Context) (entity.Tasks, error) {
				if tt.tasks != nil {
					return tt.tasks, nil
				}
				return nil, errors.New("error for mock")
			}

			sut := TaskList{Service: moq}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}

}
