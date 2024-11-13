package data

import (
	"context"
	"database/sql"
)

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	Createtimestamp string `json:"createtimestamp"`
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
		user.Username,
		user.Password,
		user.Email,
	).Scan(&user.Id, &user.Createtimestamp)

	return err
}
