package dto

type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NoteResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}
