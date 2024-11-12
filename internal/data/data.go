package data

import (
	"context"
	"database/sql"
)

type Database struct {
	Posts interface {
		Create(context.Context, *Post) error
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

func NewPostgresDB(db *sql.DB) Database {
	return Database{
		Posts: &PostgresPosts{db},
		Users: &PostgresUsers{db},
	}
}
