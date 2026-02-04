package sessions

import (
	"context"
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Create(ctx context.Context, token string, userID int64, expiresAt time.Time) error {
	const q = `
		INSERT INTO sessions (token, user_id, expires_at)
		VALUES ($1, $2, $3);
	`
	_, err := r.db.ExecContext(ctx, q, token, userID, expiresAt)
	return err
}
