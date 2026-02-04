# Gopher Notes

Мини-сервис заметок на Go: регистрация, логин по паролю, токены-сессии, CRUD заметок.
Хранение данных — PostgreSQL.

## Requirements
- Go
- Docker + Docker Compose
- migrate (golang-migrate)

## Quick start

1) Запустить PostgreSQL:
```bash
make up
```
2) Установить переменную окружения для подключения к БД:
```bash
export DATABASE_URL='postgres://app:app@localhost:5432/gopher_notes?sslmode=disable'
```
3) Применить миграции:
```bash
make migrate-up
```
4) Запустить сервер:
```bash
make run
```
Server будет доступен на `http://localhost:8080`.

## API

### Health
```bash
curl -i http://localhost:8080/health
```

### Signup
```bash
curl -i -X POST http://localhost:8080/signup \
  -H 'Content-Type: application/json' \
  -d '{"email":"a@b.com","password":"123"}'
```

### Login
```bash
curl -i -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"a@b.com","password":"123"}'
```

### Me
```bash
curl -i http://localhost:8080/me \
  -H "Authorization: Bearer <TOKEN>"
```

### Notes

Сreate:

```bash
curl -i -X POST http://localhost:8080/notes \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"title":"t1","body":"b1"}'
```

List:

```bash
curl -i http://localhost:8080/notes \
  -H "Authorization: Bearer <TOKEN>"
```

Get by id:

```bash
curl -i http://localhost:8080/notes/1 \
  -H "Authorization: Bearer <TOKEN>"
```

Delete:

```bash
curl -i -X DELETE http://localhost:8080/notes/1 \
  -H "Authorization: Bearer <TOKEN>"
```

## Project structure
-	`cmd/api` — entrypoint
-	`internal/app` — wiring (db, http server, graceful shutdown)
-	`internal/http` — handlers, middleware, router, dto
-	`internal/repository` — работа с PostgreSQL (users/sessions/notes)
-	`migrations` — миграции базы данных
