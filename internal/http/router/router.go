package router

import (
	"gopher-notes/internal/http/handlers"
	"gopher-notes/internal/http/middleware"
	"gopher-notes/internal/repository/notes"
	"gopher-notes/internal/repository/sessions"
	"gopher-notes/internal/repository/users"
	"net/http"
)

func New(usersRepo *users.Repo, sessionsRepo *sessions.Repo, notesRepo *notes.Repo) http.Handler {
	me := middleware.Auth(sessionsRepo)(handlers.Me(usersRepo))
	notesHandler := middleware.Auth(sessionsRepo)(handlers.Notes(notesRepo))
	noteByid := middleware.Auth(sessionsRepo)(handlers.NoteByID(notesRepo))
	mux := http.NewServeMux()
	mux.Handle("/notes/", noteByid)
	mux.Handle("/notes", notesHandler)
	mux.Handle("/me", me)
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/signup", handlers.Signup(usersRepo))
	mux.HandleFunc("/login", handlers.Login(usersRepo, sessionsRepo))
	return middleware.RequestLogger(mux)
}
