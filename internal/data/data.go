package data

import (
	"context"
	"database/sql"
)

type Database struct {
	Posts interface {
		Create(context.Context, *Post) error
		Get(context.Context, int64) (*Post, error)
	}

	Users interface {
		Create(context.Context, *User) error
	}

	Comment interface {
		Get(context.Context, int64) ([]Comment, error)
	}
}

func NewPostgresDB(db *sql.DB) Database {
	return Database{
		Posts:   &PostgresPosts{db},
		Users:   &PostgresUsers{db},
		Comment: &PostgresComments{db},
	}
}
