package main

import (
	"errors"
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

	comments, err := app.database.Comment.Get(ctx, postID)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	post.Comments = comments

	err = writeJSON(w, http.StatusCreated, post)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) DeletePost(w http.ResponseWriter, r *http.Request) {

	idparam := chi.URLParam(r, "postID")

	postID, err := strconv.ParseInt(idparam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	err = app.database.Posts.Delete(ctx, postID)

	if err != nil {
		app.NotFoundError(w, r, err)
	}

	w.WriteHeader(http.StatusNoContent)
}

type UpdatePostPayload struct {
	Title   string `json:"title" validate:"required,max=100"`
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) PatchPost(w http.ResponseWriter, r *http.Request) {

	idparam := chi.URLParam(r, "postID")

	postID, err := strconv.ParseInt(idparam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	var payload UpdatePostPayload

	err = readJSON(w, r, &payload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	err = Validate.Struct(payload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	if payload.Content == "" || payload.Title == "" {
		app.internalServerError(w, r, errors.New("title or content cant be empty"))
	}

	err = app.database.Posts.Update(ctx, postID, payload.Title, payload.Content)

	if err != nil {
		app.NotFoundError(w, r, err)
		return
	}
}

type CommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) AddComment(w http.ResponseWriter, r *http.Request) {

	var commentpayload CommentPayload

	err := readJSON(w, r, &commentpayload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	err = Validate.Struct(commentpayload)

	if err != nil {
		app.BadRequestError(w, r, err)
		return
	}

	idparam := chi.URLParam(r, "postID")

	postID, err := strconv.ParseInt(idparam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	userid := 1

	comment := &data.Comment{
		Content: commentpayload.Content,
		Postid:  postID,
		Userid:  int64(userid),
	}

	ctx := r.Context()

	err = app.database.Comment.Create(ctx, comment)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, comment)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
