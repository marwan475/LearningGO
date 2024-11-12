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

	// run the server and log any errors
	log.Fatal(app.run())
}
