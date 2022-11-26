package handler

import (
	"encoding/json"
	"net/http"
	"shmz_book/entity"

	"github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := ru.Validator.Struct(&body); err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	u, err := ru.Service.RegisterUser(ctx, body.Name, body.Password, body.Role)
	if err != nil {
		ResponseJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	ResponseJSON(ctx, w, rsp, http.StatusOK)
}
