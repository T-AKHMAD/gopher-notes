.PHONY: up down db-psql migrate-up migrate-down run test

up:
	docker compose up -d

down:
	docker compose down

db-psql:
	docker exec -it gopher-notes-db psql -U app -d gopher_notes

migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down 1

run:
	HTTP_ADDR=:8080 DATABASE_URL="$$DATABASE_URL" go run ./cmd/api

test:
	go test ./...