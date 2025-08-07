# GoProjects
# My First Go API

A simple todo API built with Go (Golang) using only the standard library.

## Features
- `GET /todo` – get all todos
- `POST /todo` – add a new todo
- `DELETE /todo/{id}` – delete by index

## Structure
- `internal/todo` – business logic
- `internal/transport` – HTTP handlers
- `main.go` – entry point

## Run
```bash
go run main.go
