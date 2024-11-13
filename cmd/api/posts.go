package main

import (
	"net/http"

	"github.com/marwan475/LearningGO/internal/data"
)

type PostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) CreatePost(w http.ResponseWriter, r *http.Request) {

	var postpayload PostPayload

	err := readJSON(w, r, &postpayload)

	if err != nil {
		writeJSONerror(w, http.StatusBadRequest, err.Error())
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
		writeJSONerror(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = writeJSON(w, http.StatusCreated, post)

	if err != nil {
		writeJSONerror(w, http.StatusInternalServerError, err.Error())
		return
	}
}
