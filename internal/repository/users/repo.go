package users

import (
	"context"
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateUser(ctx context.Context, email, passwordHash string) (int64, error) {
	const q = `
	INSERT INTO users (email, password_hash)
	VALUES ($1, $2)
	RETURNING id;
	`
	var id int64
	err := r.db.QueryRowContext(ctx, q, email, passwordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
