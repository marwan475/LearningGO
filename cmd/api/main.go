package main

import "log"

func main() {

	// app config
	cfg := config{
		addr: ":8080",
	}

	// main app
	app := &application{
		config: cfg,
	}

	// api route handler
	mux := app.mount()

	// run the server and log any errors
	log.Fatal(app.run(mux))
}
