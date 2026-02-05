![CI](https://github.com/T-AKHMAD/gopher-notes/actions/workflows/ci.yml/badge.svg)

# Gopher Notes

Разработал REST API сервис заметок на Go: signup/login, bearer-token sessions, защищённые эндпоинты, PostgreSQL, миграции и Docker Compose. Реализовал middleware для логирования и авторизации, graceful shutdown.

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
или положить в окружение любым способом.

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

### Logout
```bash
curl -i -X POST http://localhost:8080/logout \
  -H "Authorization: Bearer <TOKEN>"
```

### Me
```bash
curl -i http://localhost:8080/me \
  -H "Authorization: Bearer <TOKEN>"
```

### Notes

Create:

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
- `cmd/api` — entrypoint
- `internal/app` — wiring (db, http server, graceful shutdown)
- `internal/http` — handlers, middleware, router, dto
- `internal/repository` — работа с PostgreSQL (users/sessions/notes)
- `migrations` — миграции базы данных

Все ошибки возвращаются в JSON, например:

```json
{"error":"..."}
```

## Future improvements

- Config struct + env validation
- Cleanup expired sessions (cron / background job)
- Pagination for notes list