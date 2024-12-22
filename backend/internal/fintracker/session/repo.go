package session

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrTokenNotFound = errors.New("token not found")
)

type Session struct {
	UserID  uint64
	Token   string
	IsValid bool
}

type Repository interface {
	Get(ctx context.Context, token string) (*Session, error)
	Create(ctx context.Context, s Session) error
	Invalidate(ctx context.Context, userID uint64) error
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

	var sess Session
	err := row.Scan(&sess.UserID, &sess.Token, &sess.IsValid)
	if err != nil {
		return nil, err
	}

	return &sess, nil
}

func (r *repo) Create(ctx context.Context, s Session) error {
	_, err := r.db.ExecContext(ctx, createSessionQuery, s.Token, s.UserID, true)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Invalidate(ctx context.Context, userID uint64) error {
	res, err := r.db.ExecContext(ctx, invalidateSessionQuery, userID)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if aff == 0 {
		return ErrTokenNotFound
	}

	return nil
}
