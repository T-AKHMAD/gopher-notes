package handlers

import (
	"encoding/json"
	"gopher-notes/internal/repository/users"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(usersRepo *users.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		var req signupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}
		if req.Email == "" || req.Password == "" {
			writeError(w, http.StatusBadRequest, "email and password are required")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}

		id, err := usersRepo.CreateUser(r.Context(), req.Email, string(hash))
		if err != nil {
			msg := err.Error()
			if strings.Contains(msg, "duplicate key") || strings.Contains(msg, "users_email_key") {
				writeError(w, http.StatusConflict, "email already exists")
				return
			}
			writeError(w, http.StatusInternalServerError, "internal error")
			return
		}

		_ = writeJSON(w, http.StatusCreated, map[string]any{
			"id": id,
		})
	}
}
