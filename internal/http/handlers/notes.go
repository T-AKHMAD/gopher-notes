package handlers

import (
	"encoding/json"
	"gopher-notes/internal/http/dto"
	"gopher-notes/internal/http/middleware"
	"gopher-notes/internal/repository/notes"
	"net/http"
	"time"
)

func Notes(notesRepo *notes.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		switch r.Method {
		case http.MethodPost:
			var req dto.CreateNoteRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeError(w, http.StatusBadRequest, "invalid json")
				return
			}
			if req.Title == "" || req.Body == "" {
				writeError(w, http.StatusBadRequest, "title and body are required")
				return
			}
			id, err := notesRepo.Create(r.Context(), userID, req.Title, req.Body)
			if err != nil {
				writeError(w, http.StatusInternalServerError, "internal error")
				return
			}
			_ = writeJSON(w, http.StatusCreated, map[string]any{"id": id})
			return
		case http.MethodGet:
			list, err := notesRepo.ListByUserID(r.Context(), userID)
			if err != nil {
				writeError(w, http.StatusInternalServerError, "internal error")
				return
			}
			resp := make([]dto.NoteResponse, 0, len(list))
			for _, n := range list {
				resp = append(resp, dto.NoteResponse{
					ID:        n.ID,
					Title:     n.Title,
					Body:      n.Body,
					CreatedAt: n.CreatedAt.Format(time.RFC3339),
				})
			}
			_ = writeJSON(w, http.StatusOK, map[string]any{"notes": resp})
			return
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
	}
}
