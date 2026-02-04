package notes

import "context"

func (r *Repo) GetByID(ctx context.Context, userID, id int64) (*Note, error) {
	const q = `
		SELECT id, user_id, title, body, created_at
		FROM notes
		WHERE id = $1 AND user_id = $2;
	`

	var n Note
	err := r.db.QueryRowContext(ctx, q, id, userID).Scan(
		&n.ID, &n.UserID, &n.Title, &n.Body, &n.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
