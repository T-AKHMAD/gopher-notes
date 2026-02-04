package sessions

import (
	"context"
	"time"
)

type Session struct {
	Token     string
	UserID    int64
	ExpiresAt time.Time
}

func (r *Repo) GetByToken(ctx context.Context, token string) (*Session, error) {
	const q = `
		SELECT token, user_id, expires_at
		FROM sessions
		WHERE token = $1;
	`

	var s Session
	err := r.db.QueryRowContext(ctx, q, token).Scan(&s.Token, &s.UserID, &s.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
