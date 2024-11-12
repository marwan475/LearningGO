package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// main api
type application struct {
	config config
}

// config
type config struct {
	addr string
}

// mount the applications router
func (app *application) mount() http.Handler {

	// http router using chi
	r := chi.NewRouter()

	// midleware for routter (currently logs requests and responses)
	r.Use(middleware.Logger)

	// routes
	r.Get("/v1/health", app.CheckHealth)

	return r
}

// run method for application
func (app *application) run(mux http.Handler) error {

	// setting up http server on addr
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // if server takes more the 30s to write then time out
		ReadTimeout:  time.Second * 10, //  if client takes more then 10s to read our response the time out
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started on %v \n", app.config.addr)

	// start server and return any errors
	return server.ListenAndServe()
}
