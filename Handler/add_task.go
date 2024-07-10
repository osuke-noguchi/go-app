package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/osuke-noguchi/go-app/entity"
	"github.com/osuke-noguchi/go-app/store"
)

type AddTask struct {
	Store *store.TaskStore
	Validate *validator.Validate
}

func (at *AddTask) SevedHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err := nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title: b.Title,
		Status: entity.TaskStatusTodo,
		Created: time.Now(),
	}
	id, err := store.Tasks.Add(t)
	if err := nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{ID: id}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}