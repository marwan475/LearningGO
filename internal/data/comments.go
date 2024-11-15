package data

import (
	"context"
	"database/sql"
)

type Comment struct {
	Id              int64  `json:"id"`
	Postid          int64  `json:"postid"`
	Userid          int64  `json:"userid"`
	Content         string `json:"content"`
	Createtimestamp string `json:"createtimestamp"`
	User            User   `json:"user"`
}

type PostgresComments struct {
	db *sql.DB
}

func (s *PostgresComments) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments (postid, userid, content)
		VALUES ($1, $2, $3)
		`

	err := s.db.QueryRowContext(ctx, query,
		comment.Postid,
		comment.Userid,
		comment.Content,
	).Err()

	return err

}

func (s *PostgresComments) Get(ctx context.Context, id int64) ([]Comment, error) {
	query := `
		SELECT c.postid, c.userid, c.content, c.createtimestamp FROM comments c
		WHERE c.postid = $1
		ORDER BY c.createtimestamp DESC;
	`
	rows, err := s.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.Postid, &c.Userid, &c.Content, &c.Createtimestamp)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil

}
