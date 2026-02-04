package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"gopher-notes/internal/repository/sessions"
	"gopher-notes/internal/repository/users"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(usersRepo *users.Repo, sessionsRepo *sessions.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		if req.Email == "" || req.Password == "" {
			writeError(w, http.StatusBadRequest, "email and password are required")
			return
		}
		u, err := usersRepo.GetByEmail(r.Context(), req.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				writeError(w, http.StatusUnauthorized, "invalid credentials")
				return
			}
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
			writeError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}

		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}
		token := hex.EncodeToString(b)

		expiresAt := time.Now().Add(24 * time.Hour)

		if err := sessionsRepo.Create(r.Context(), token, u.ID, expiresAt); err != nil {
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}

		_ = writeJSON(w, http.StatusOK, map[string]string{
			"token":      token,
			"expires_at": expiresAt.Format(time.RFC3339),
		})
	}
}
