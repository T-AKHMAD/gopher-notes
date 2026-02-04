package main

import (
	"gopher-notes/internal/app"
	"log"
	"os"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	a := app.New(addr, dbURL)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
