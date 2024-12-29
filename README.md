# Hexa-Go-CRUD

A CRUD application built with Go using Hexagonal Architecture (Ports and Adapters pattern) and PostgreSQL.

## Technologies

- Go 1.24
- PostgreSQL 17.2
- Docker & Docker Compose

## Project Structure

```
.
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   │   └── repository/
│   └── domain/
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Prerequisites

- Docker
- Docker Compose
- Go 1.24 (for local development)

## Getting Started

1. Clone the repository:
```bash
cd hexa-go-crud
```

2. Start the application:
```bash
docker compose up --build
```

The application will be available at `http://localhost:8080`


## API Endpoints

- `GET /api/something` - List all items
- `POST /api/something` - Create a new item
- `GET /api/something/:id` - Get an item by ID
- `PUT /api/something/:id` - Update an item
- `DELETE /api/something/:id` - Delete an item

## Environment Variables

- `DB_HOST` - PostgreSQL host
- `DB_PORT` - PostgreSQL port
- `POSTGRES_USER` - PostgreSQL username
- `POSTGRES_PASSWORD` - PostgreSQL password
- `DB_NAME` - Database name

## License

[MIT License](LICENSE)
