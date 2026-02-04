package notes

import "context"

func (r *Repo) Create(ctx context.Context, userID int64, title, body string) (int64, error) {
	const q = `
		INSERT INTO notes (user_id, title, body)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id int64
	err := r.db.QueryRowContext(ctx, q, userID, title, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
