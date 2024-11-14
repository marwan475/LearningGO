package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	Id              int64     `json:"id"`
	Content         string    `json:"content"`
	Title           string    `json:"title"`
	Userid          int64     `json:"userid"`
	Createtimestamp string    `json:"createtimestamp"`
	Updatetimestap  string    `json:"updatetimestamp"`
	Tags            []string  `json:"tags"`
	Comments        []Comment `json:"comments"`
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
		post.Content,
		post.Title,
		post.Userid,
		pq.Array(post.Tags),
	).Scan(&post.Id, &post.Createtimestamp, &post.Updatetimestap)

	return err
}

func (s *PostgresPosts) Get(ctx context.Context, id int64) (*Post, error) {
	query := `
	SELECT id, userid, title, content, createtimestamp, updatetimestamp, tags
	FROM posts
	WHERE id = $1
	`

	var post Post

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.Id,
		&post.Userid,
		&post.Title,
		&post.Createtimestamp,
		&post.Updatetimestap,
		&post.Tags,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &post, nil

}
