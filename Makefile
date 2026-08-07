APP_NAME := student-management

# Default target
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo "  run         run the API server locally"
	@echo "  migrate     run auto migrations and seed data, then exit"
	@echo "  swagger     regenerate Swagger docs"
	@echo "  test        run all unit tests"
	@echo "  lint        run golangci-lint"
	@echo "  build       build the binary into ./bin"
	@echo "  docker-up   build and start docker compose services"
	@echo "  docker-down stop docker compose services"
	@echo "  k8s         deploy to Minikube"

.PHONY: run
run:
	go run ./cmd/server

.PHONY: migrate
migrate:
	go run ./cmd/server -migrate

.PHONY: swagger
swagger:
	swag init -g cmd/server/main.go -o docs

.PHONY: test
test:
	go test ./... -count=1

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: docker-up
docker-up:
	docker compose up -d --build

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: k8s
k8s:
	./scripts/minikube-setup.sh