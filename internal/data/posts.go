package data

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	id              int64    `json:"id"`
	content         string   `json:"content"`
	title           string   `json:"title"`
	userid          int64    `json:"userid"`
	createtimestamp string   `json:"createtimestamp"`
	updatetimestap  string   `json:"updatetimestamp"`
	tags            []string `json:"tags"`
}

type PostgresPosts struct {
	db *sql.DB
}

func (s *PostgresPosts) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, userid, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, createtimestamp, updatetimestamp
		`
	err := s.db.QueryRowContext(ctx, query,
		post.content,
		post.title,
		post.userid,
		pq.Array(post.tags),
	).Scan(&post.id, &post.createtimestamp, &post.updatetimestap)

	return err
}
