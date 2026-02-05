package middleware

import (
	"context"
	"database/sql"
	"errors"
	"gopher-notes/internal/http/response"
	"gopher-notes/internal/repository/sessions"
	"net/http"
	"strings"
	"time"
)

type ctxKey string

const userIDKey ctxKey = "user_id"

func UserIDFromContext(ctx context.Context) (int64, bool) {
	v := ctx.Value(userIDKey)
	id, ok := v.(int64)
	return id, ok
}

func Auth(sessionsRepo *sessions.Repo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := bearerToken(r)
			if !ok {
				response.WriteError(w, http.StatusUnauthorized, "missing or invalid authorization")
				return
			}
			sess, err := sessionsRepo.GetByToken(r.Context(), token)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					response.WriteError(w, http.StatusUnauthorized, "invalid token")
					return
				}
				response.WriteError(w, http.StatusInternalServerError, "internal error")
				return
			}

			if time.Now().After(sess.ExpiresAt) {
				response.WriteError(w, http.StatusUnauthorized, "token expired")
				return
			}
			ctx := context.WithValue(r.Context(), userIDKey, sess.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func bearerToken(r *http.Request) (string, bool) {
	h := r.Header.Get("Authorization")
	if h == "" {
		return "", false
	}
	const prefix = "Bearer "
	if !strings.HasPrefix(h, prefix) {
		return "", false
	}

	token := strings.TrimSpace(strings.TrimPrefix(h, prefix))
	if token == "" {
		return "", false
	}
	return token, true
}

func BearerToken(r *http.Request) (string, bool) {
	return bearerToken(r)
}
