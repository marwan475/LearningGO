package data

import (
	"context"
	"database/sql"
)

type Database struct {
	Posts interface {
		Create(context.Context, *Post) error
		Get(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, int64, string, string) error
	}

	Users interface {
		Create(context.Context, *User) error
	}

	Comment interface {
		Create(context.Context, *Comment) error
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
