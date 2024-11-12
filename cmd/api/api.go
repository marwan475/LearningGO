package main

import "net/http"

// main api
type application struct {
	config config
}

// config
type config struct {
	addr string
}

// method for application
func (app *application) run() error {

	// http router
	mux := http.NewServeMux()

	// setting up http server on addr
	server := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}

	// start server and return any errors
	return server.ListenAndServe()
}
