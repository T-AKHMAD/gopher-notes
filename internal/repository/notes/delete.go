package notes

import (
	"context"
)

func (r *Repo) Delete(ctx context.Context, userID, id int64) (bool, error) {
	const q = `
		DELETE FROM notes
		WHERE id = $1 AND user_id = $2;
	`

	res, err := r.db.ExecContext(ctx, q, id, userID)
	if err != nil {
		return false, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}
