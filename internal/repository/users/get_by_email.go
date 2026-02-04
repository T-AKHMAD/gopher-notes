package users

import (
	"context"
)

type User struct {
	ID           int64
	Email        string
	PasswordHash string
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (*User, error) {
	const q = `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1;
	`

	var u User
	err := r.db.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
