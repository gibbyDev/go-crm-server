# Simple Makefile for a Go project

# Build the application
all: build 

build:
	@echo "Building..."
	@go build -o /tmp/main cmd/api/main.go

tailwind-watch:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

tailwind-build:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

templ-watch:
	templ generate --watch

templ-generate:
	templ generate

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f /tmp/main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching..."; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
			echo "Watching..."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Build and run the auth service
auth-build:
	@go build -o /tmp/auth cmd/auth/main.go

auth-run:
	@/tmp/auth

# Build and run the billing service
billing-build:
	@go build -o billing cmd/billing/main.go

billing-run:
	@go run cmd/billing/main.go

.PHONY: all build run clean watch docker-run docker-down auth-build auth-run billing-build billing-run tailwind-watch tailwind-build templ-watch templ-generate
