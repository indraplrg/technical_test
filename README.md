# Student Management System

Production-ready REST API for managing **Jurusan** (departments) and **Mahasiswa** (students),
built as a Senior Backend Engineer technical test.

## Tech Stack

- **Go 1.24+**
- **Gin Framework** – HTTP router / middleware
- **PostgreSQL** – relational database
- **GORM** – ORM with AutoMigrate and Soft Delete
- **Swaggo** – Swagger / OpenAPI documentation
- **Docker & Docker Compose** – containerization
- **Minikube** – local Kubernetes deployment
- **MVC Architecture** – controller / service / repository layers

## Features

### CRUD

- CRUD `Jurusan` (create, list, get by id, update, delete)
- CRUD `Mahasiswa` (create, list, get by id, update, delete)
- Search (`search`, `nim`), filter (`id_jurusan`), sorting (`sort_by`, `sort_order`) and pagination (`page`, `limit`) on Mahasiswa list

### Additional Requirements

- Soft Delete (`deleted_at`) + audit fields (`created_at`, `updated_at`)
- Request ID middleware (`X-Request-ID`)
- CORS middleware
- Recovery middleware with consistent JSON errors
- Health check endpoint `GET /health`
- API versioning under `/api/v1`
- Environment configuration via `.env`
- Seed data (5 jurusan on first boot)
- Makefile (`make run`, `make migrate`, `make swagger`, `make test`)
- golangci-lint ready (`.golangci.yml`)
- Testable repository & service architecture
- RESTful naming conventions
- UTC timezone used internally
- Dependency injection (constructor-based)
- Graceful shutdown
- Structured logging (`log/slog` JSON)
- Context propagation across layers
- Exports: CSV, Excel, PDF, JSON

## Folder Structure

```
cmd/
  server/            # application entrypoint
internal/
  config/            # environment configuration loader
  controller/        # thin HTTP handlers
  service/           # business logic + domain errors
  repository/        # data access (GORM)
  model/             # GORM models
  dto/               # request/response payloads
  middleware/        # request id, CORS, recovery
  validator/         # request validation + custom rules
  routes/            # router + dependency wiring
  helper/            # shared helpers
  response/          # consistent JSON response envelope
  database/          # connection, migration, seed
docs/                # generated Swagger docs
migrations/          # reference SQL schema
scripts/             # tooling scripts (minikube setup)
k8s/                 # Kubernetes manifests
```

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL 14+ (or Docker)
- [swag](https://github.com/swaggo/swag) for regenerating docs
- (optional) Docker / Minikube / golangci-lint

### 1. Configure environment

```bash
cp .env.example .env
```

Adjust values in `.env` as needed.

### 2. Run locally

```bash
make migrate   # create schema + seed data
make run       # start server on :8080
```

Or directly:

```bash
go run ./cmd/server -migrate
go run ./cmd/server
```

### 3. Swagger UI

```bash
make swagger
```

Open `http://localhost:8080/swagger/index.html`.

## API

Base path: `/api/v1`

### Health

| Method | Path       | Description        |
| ------ | ---------- | ------------------ |
| GET    | `/health`  | Service health     |
| GET    | `/swagger/*any` | Swagger UI    |

### Jurusan

| Method | Path                  | Description        |
| ------ | --------------------- | ------------------ |
| POST   | `/api/v1/jurusan`     | Create jurusan     |
| GET    | `/api/v1/jurusan`     | List all jurusan   |
| GET    | `/api/v1/jurusan/:id` | Get jurusan by id  |
| PUT    | `/api/v1/jurusan/:id` | Update jurusan     |
| DELETE | `/api/v1/jurusan/:id` | Delete jurusan     |

### Mahasiswa

| Method | Path                          | Description                          |
| ------ | ----------------------------- | ------------------------------------ |
| POST   | `/api/v1/mahasiswa`           | Create mahasiswa                     |
| GET    | `/api/v1/mahasiswa`           | List with search/filter/sort/paginate |
| GET    | `/api/v1/mahasiswa/:id`       | Get mahasiswa by id                  |
| PUT    | `/api/v1/mahasiswa/:id`       | Update mahasiswa                     |
| DELETE | `/api/v1/mahasiswa/:id`       | Delete mahasiswa (soft)              |
| GET    | `/api/v1/mahasiswa/export/csv`   | Export CSV                        |
| GET    | `/api/v1/mahasiswa/export/excel` | Export Excel                      |
| GET    | `/api/v1/mahasiswa/export/pdf`   | Export PDF                        |
| GET    | `/api/v1/mahasiswa/export/json`  | Export JSON                       |

### List query parameters (Mahasiswa)

| Param        | Type   | Description                                 |
| ------------ | ------ | ------------------------------------------- |
| `search`     | string | partial match on `nama` or `nim`            |
| `nim`        | string | exact match on `nim`                        |
| `id_jurusan` | int    | filter by department                        |
| `sort_by`    | string | `nama`, `umur`, `nim`, `tanggal_lahir`, `created_at` |
| `sort_order` | string | `asc` / `desc`                              |
| `page`       | int    | page number (default 1)                     |
| `limit`      | int    | page size (default 10, max 100)             |

### Example requests

Create jurusan:

```bash
curl -X POST http://localhost:8080/api/v1/jurusan \
  -H "Content-Type: application/json" \
  -d '{"nama_jurusan":"Teknik Informatika","fakultas":"Fakultas Ilmu Komputer","jenjang":"S1"}'
```

Create mahasiswa:

```bash
curl -X POST http://localhost:8080/api/v1/mahasiswa \
  -H "Content-Type: application/json" \
  -d '{"nama":"Budi Santoso","umur":21,"nim":"TI-2026-001","tanggal_lahir":"2005-03-15","alamat":"Jakarta","id_jurusan":1}'
```

List with search and pagination:

```bash
curl "http://localhost:8080/api/v1/mahasiswa?search=budi&sort_by=created_at&sort_order=desc&page=1&limit=10"
```

Export CSV:

```bash
curl -OJ http://localhost:8080/api/v1/mahasiswa/export/csv
```

### Response format

Success:

```json
{
  "success": true,
  "message": "jurusan created",
  "data": { "id_jurusan": 1, "nama_jurusan": "Teknik Informatika" }
}
```

Paginated success:

```json
{
  "success": true,
  "message": "mahasiswa fetched successfully",
  "data": {
    "items": [],
    "pagination": { "page": 1, "limit": 10, "total": 0, "total_pages": 0 }
  }
}
```

Error:

```json
{
  "success": false,
  "message": "validation error",
  "errors": ["nama is required"]
}
```

### HTTP status codes

| Code | Meaning                                   |
| ---- | ----------------------------------------- |
| 200  | OK / updated / deleted                    |
| 201  | Created                                   |
| 400  | Bad request / validation error            |
| 404  | Resource not found                        |
| 409  | Conflict (duplicate nim / name, related records) |
| 500  | Internal server error                     |

## Validation Rules

**Mahasiswa**

- `nama` required
- `umur` required, > 0
- `nim` required, unique
- `tanggal_lahir` required, format `YYYY-MM-DD`
- `alamat` required
- `id_jurusan` must exist

**Jurusan**

- `nama_jurusan` required
- `fakultas` required
- `jenjang` required

## Docker

```bash
docker compose up -d --build   # starts postgres + backend
docker compose down            # stop
```

Services: `backend` (port 8080) and `postgres` (port 5432). Migrations and seed run automatically on startup.

## Kubernetes (Minikube)

```bash
./scripts/minikube-setup.sh
```

or

```bash
make k8s
```

The script starts Minikube, builds the image into the Minikube Docker daemon, applies all manifests in `k8s/` and prints the service URL. The backend is exposed on NodePort `30080`.

## Makefile

```bash
make run        # run API server
make migrate    # run migrations + seed
make swagger    # regenerate Swagger docs
make test       # run unit tests
make lint       # golangci-lint
make docker-up  # docker compose up -d --build
make docker-down
make k8s        # deploy to Minikube
```

## Testing

```bash
go test ./... -count=1
```

The project is structured so repository and service layers can be tested with mocks:

- repositories depend only on `*gorm.DB` (or an interface)
- services depend on repository interfaces

## Linting

```bash
golangci-lint run ./...
```

Configuration lives in `.golangci.yml`.
