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
# Convenience target for reset
.PHONY: reset
reset: clean up
