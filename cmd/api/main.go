package main

import (
	"log"

	"github.com/marwan475/LearningGO/internal/data"
	"github.com/marwan475/LearningGO/internal/db"
	"github.com/marwan475/LearningGO/internal/env"
)

func main() {

	// app config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DBADDR", "postgres://admin:adminpassword@localhost/testdb?sslmode=disable"),
			maxCon:      env.GetInt("DBCON", 30),
			maxIdle:     env.GetInt("DBIDLE", 30),
			maxIdletime: env.GetString("DBIDLETIME", "15m"),
		},
	}

	database, err := db.New(cfg.db.addr, cfg.db.maxCon, cfg.db.maxIdle, cfg.db.maxIdletime)

	if err != nil {
		log.Panic(err)
	}

	defer database.Close()
	log.Println("Database Connection Established")

	// Data base interface
	datab := data.NewPostgresDB(database)

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
