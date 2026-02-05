package handlers

import (
	"gopher-notes/internal/http/middleware"
	"gopher-notes/internal/repository/sessions"
	"net/http"
)

func Logout(sessionsRepo *sessions.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return 
		}
		token, ok := middleware.BearerToken(r)
		if !ok {
			writeError(w, http.StatusUnauthorized, "missing or invalid authorization")
			return 
		}
		deleted, err := sessionsRepo.Delete(r.Context(), token)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "internal error")
			return 
		}
		if !deleted {
			writeError(w, http.StatusUnauthorized, "invalid token")
			return 
		}
		w.WriteHeader(http.StatusNoContent)
	}
}