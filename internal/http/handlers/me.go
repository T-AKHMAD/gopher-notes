package handlers

import (
	"database/sql"
	"errors"
	"gopher-notes/internal/http/middleware"
	"gopher-notes/internal/repository/users"
	"net/http"
)

func Me(usersRepo *users.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := middleware.UserIDFromContext(r.Context())
		if !ok {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		u, err := usersRepo.GetByID(r.Context(), userID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				writeError(w, http.StatusUnauthorized, "unauthorized")
				return
			}
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}
		_ = writeJSON(w, http.StatusOK, map[string]any{
			"id":    u.ID,
			"email": u.Email,
		})
	}
}
