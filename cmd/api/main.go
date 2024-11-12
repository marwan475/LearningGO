package main

import (
	"log"

	"github.com/marwan475/LearningGO/internal/data"
	"github.com/marwan475/LearningGO/internal/env"
)

func main() {

	// app config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	// Data base interface
	datab := data.NewPostgresDB(nil)

	// main app
	app := &application{
		config:   cfg,
		database: datab,
	}

	// api route handler
	mux := app.mount()

	// run the server and log any errors
	log.Fatal(app.run(mux))
}
