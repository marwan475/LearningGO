package main

import (
	"net/http"

	"github.com/marwan475/LearningGO/internal/data"
)

type UserPayload struct {
	Username string `json:"username" validate:"required,max=30"`
	Email    string `json:"email" validate:"required,max=50"`
	Password string `json:"-"`
}

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {

	var userpaylod UserPayload

	err := readJSON(w, r, &userpaylod)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = Validate.Struct(userpaylod)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	user := &data.User{
		Username: userpaylod.Username,
		Email:    userpaylod.Email,
		Password: userpaylod.Password,
	}

	ctx := r.Context()

	err = app.database.Users.Create(ctx, user)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, user)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
