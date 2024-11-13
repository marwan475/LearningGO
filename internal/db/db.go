package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxCon, maxIdle int, maxIdletime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxCon)
	db.SetMaxIdleConns(maxIdle)

	duration, err := time.ParseDuration(maxIdletime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
