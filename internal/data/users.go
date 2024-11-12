package data

import (
	"context"
	"database/sql"
)

type User struct {
	id              int64  `json: "id"`
	username        string `json: "username"`
	email           string `json: "email"`
	password        string `json: "-"`
	createtimestamp string `json: "createtimestamp"`
}

type PostgresUsers struct {
	db *sql.DB
}

func (s *PostgresUsers) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES ($1, $2, $3) RETURNING id, createtimestamp
	`
	err := s.db.QueryRowContext(ctx, query,
		user.username,
		user.password,
		user.email,
	).Scan(&user.id, &user.createtimestamp)

	return err
}
