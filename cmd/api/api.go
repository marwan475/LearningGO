package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marwan475/LearningGO/internal/data"
)

// main api
type application struct {
	config   config
	database data.Database
}

// config
type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr        string
	maxCon      int
	maxIdle     int
	maxIdletime string
}

// mount the applications router
func (app *application) mount() http.Handler {

	// http router using chi
	r := chi.NewRouter()

	// midleware
	r.Use(middleware.RequestID) // request id
	r.Use(middleware.RealIP)    // real ip
	r.Use(middleware.Recoverer) // recovers from panics
	r.Use(middleware.Logger)    // logs requests
	r.Use(middleware.Timeout(60 * time.Second))

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
