package sessions

import "context"

func (r *Repo) Delete(ctx context.Context, token string) (bool, error) {
	const q = `
		DELETE FROM sessions
		WHERE token = $1;
	`

	res, err := r.db.ExecContext(ctx, q, token)
	if err != nil {
		return false, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return n > 0, nil
}