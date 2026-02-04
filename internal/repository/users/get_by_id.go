package users

import "context"

func (r *Repo) GetByID(ctx context.Context, id int64) (*User, error) {
	const q = `
		SELECT id, email, password_hash
		FROM users
		WHERE id = $1;
	`
	var u User
	err := r.db.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
