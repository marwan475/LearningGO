package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internel server error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONerror(w, http.StatusInternalServerError, "Server Encountered an Error")
}

func (app *application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("Bad Request error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONerror(w, http.StatusBadRequest, err.Error())
}

func (app *application) NotFoundError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("Not Found error: %s path: %s error: %s", r.Method, r.URL.Path, err.Error())

	writeJSONerror(w, http.StatusNotFound, "Not Found")
}
