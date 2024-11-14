package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marwan475/LearningGO/internal/data"
)

type PostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) CreatePost(w http.ResponseWriter, r *http.Request) {

	var postpayload PostPayload

	err := readJSON(w, r, &postpayload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	err = Validate.Struct(postpayload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	userid := 1

	post := &data.Post{
		Title:   postpayload.Title,
		Content: postpayload.Content,
		Tags:    postpayload.Tags,
		Userid:  int64(userid),
	}

	ctx := r.Context()

	err = app.database.Posts.Create(ctx, post)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, post)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) GetPost(w http.ResponseWriter, r *http.Request) {
	idparam := chi.URLParam(r, "postID")

	postID, err := strconv.ParseInt(idparam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.database.Posts.Get(ctx, postID)

	if err != nil {
		app.NotFoundError(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, post)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
