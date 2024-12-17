package session

import (
	"context"
	"database/sql"
)

type Session struct {
	UserID uint64
	Token  string
}

type Repository interface {
	Get(ctx context.Context, token string) (*Session, error)
	Create(s Session) error
	Invalidate(token string) error
}

type repo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Get(ctx context.Context, token string) (*Session, error) {
	row := r.db.QueryRowContext(ctx, getSessionQuery, token)

	var sess *Session
	err := row.Scan(&sess)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

func (r *repo) Create(s Session) error {
	_, err := r.db.Exec(createSessionQuery, s.Token, s.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Invalidate(token string) error {
	return nil
}
