package main

import (
	"log"
	"net/http"
	"time"
)

// main api
type application struct {
	config config
}

// config
type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	// http router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.CheckHealth)

	return mux
}

// method for application
func (app *application) run(mux *http.ServeMux) error {

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
