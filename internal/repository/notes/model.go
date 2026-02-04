package notes

import "time"

type Note struct {
	ID        int64
	UserID    int64
	Title     string
	Body      string
	CreatedAt time.Time
}
