package handler

import (
	"net/http"
	"testing"

	"github.com/osuke-noguchi/go-app/entity"
	"github.com/osuke-noguchi/go-app/testutil"
)

func TestAddTask(t *testing.T) {
	t.Parallel()
	type want struct {
		status int
		rspFile string
	}
	tests := map[string]struct {
		reqFile string
		want want
	}{
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status: http.StatusOK,
				rspFile: "testdata/add_task/ok_rsp.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_rsp.json.golden",
			want: want{
				status: http.StatusBadRequest,
				rspFile: "testdata/add_task/bad_req_rsp.json.golden"
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MEthodPost,
				"/tasks",
				bytes.NewRecorder(testutil.LoadFile(t, tt.reqFile)),
			)

			sut := AddTask{
				Store: &store.TaskStore{
					Tasks: map[entity.TaskID]*entity.Task{},
				},
				Validator: validator.New(),
			}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			test.AssertResponse(t,
			resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
		)
		})
	}
}