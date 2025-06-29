# Makefile for Go + PostgreSQL + sqlc MVC project

APP_NAME=golang-mvc-app

# Default target
.PHONY: all
all: generate build

# ---------------------------------------------------
# sqlc code generation
.PHONY: generate
generate:
	@echo "🚀 Generating Go code from SQL with sqlc..."
	sqlc generate

# ---------------------------------------------------
# Build
.PHONY: build
build:
	@echo "🔨 Building app with full insecure bypass..."
	GOPRIVATE=github.com GOINSECURE=github.com GOSUMDB=off go build -o $(APP_NAME) main.go

# ---------------------------------------------------
# Run
.PHONY: run
run:
	@echo "🏃 Running app..."
	go run main.go

# ---------------------------------------------------
# Docker database management
.PHONY: up
up:
	@echo "🐘 Starting PostgreSQL with Docker Compose..."
	docker-compose up -d

.PHONY: down
down:
	@echo "🧹 Stopping and removing containers..."
	docker-compose down

.PHONY: clean
clean:
	@echo "💣 Removing containers and volumes..."
	docker-compose down -v

# ---------------------------------------------------
# Wait for DB
.PHONY: wait-for-db
wait-for-db:
	@echo "⏳ Waiting for PostgreSQL to be ready..."
	@until docker exec postgres_db pg_isready -U myuser -d mydatabase > /dev/null 2>&1; do \
		echo "🔄 Waiting for db..."; \
		sleep 1; \
	done
	@echo "✅ PostgreSQL is ready."

# ---------------------------------------------------
# Convenience targets
.PHONY: reset
reset: clean up

.PHONY: dev
dev:
	@echo "🚀 Running full dev environment (reset DB, generate code, wait for DB, run app)..."
	make reset
	make wait-for-db
	make generate
	make run
