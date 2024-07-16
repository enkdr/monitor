# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# watch:
# 	@echo "watching for file changes..."
# 	@find . -name "*.go" -o -name "*.html" -o -name "*.css" -o -name "*.js" | entr -r sh -c 'kill -9 $$PID; make run & PID=$$!'


# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Create DB container
docker-pg-run:
	@if docker compose up postgres 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up postgres; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi
