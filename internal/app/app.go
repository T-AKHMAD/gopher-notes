package app

import (
	"context"
	"database/sql"
	"gopher-notes/internal/http/router"
	"gopher-notes/internal/repository/notes"
	"gopher-notes/internal/repository/sessions"
	"gopher-notes/internal/repository/users"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type App struct {
	addr  string
	dbURL string
}

func New(addr, dbURL string) *App {
	return &App{
		addr:  addr,
		dbURL: dbURL,
	}
}

func (a *App) Run() error {
	db, err := sql.Open("pgx", a.dbURL)
	if err != nil {
		return err
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		return err
	}
	defer db.Close()

	usersRepo := users.New(db)
	sessionsRepo := sessions.New(db)
	notesRepo := notes.New(db)
	srv := &http.Server{
		Addr:         a.addr,
		Handler:      router.New(usersRepo, sessionsRepo, notesRepo),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	errCh := make(chan error, 1)

	go func() {
		errCh <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			_ = srv.Close()
			return err
		}
		return nil
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}
