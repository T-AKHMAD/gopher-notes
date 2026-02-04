package notes

import "context"

func (r *Repo) ListByUserID(ctx context.Context, userID int64) ([]Note, error) {
	const q = `
		SELECT id, user_id, title, body, created_at
		FROM notes
		WHERE user_id = $1
		ORDER BY created_at DESC, id DESC;
	`
	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Note
	for rows.Next() {
		var n Note
		if err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Body, &n.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
