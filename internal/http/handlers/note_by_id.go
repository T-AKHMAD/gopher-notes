package handlers

import (
	"database/sql"
	"errors"
	"gopher-notes/internal/http/dto"
	"gopher-notes/internal/http/middleware"
	"gopher-notes/internal/repository/notes"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NoteByID(notesRepo *notes.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
		if idStr == "" || strings.Contains(idStr, "/") {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil || id <= 0 {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}
		switch r.Method {
		case http.MethodGet:

			n, err := notesRepo.GetByID(r.Context(), userID, id)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					writeError(w, http.StatusNotFound, "not found")
					return
				}
				writeError(w, http.StatusInternalServerError, "internal error")
				return
			}

			resp := dto.NoteResponse{
				ID:        n.ID,
				Title:     n.Title,
				Body:      n.Body,
				CreatedAt: n.CreatedAt.Format(time.RFC3339),
			}

			_ = writeJSON(w, http.StatusOK, resp)

			return

		case http.MethodDelete:
			deleted, err := notesRepo.Delete(r.Context(), userID, id)
			if err != nil {
				writeError(w, http.StatusInternalServerError, "internal error")
				return
			}
			if !deleted {
				writeError(w, http.StatusNotFound, "not found")
				return
			}
			return
		default:
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
	}
}
