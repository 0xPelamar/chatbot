.PHONY: start dev-db dev-db-stop test test-integration prod prod-stop build clean
# ==============================================================================
# Development Workflow
# Strategy: Run Go app locally, Run DBs in Docker
# ==============================================================================

# 1. Start the DBs (Postgres & Redis) in Docker
dev-db:
	docker compose -f docker-compose.dev.yml up -d

# 2. Stop the DBs
dev-db-stop:
	docker compose -f docker-compose.dev.yml down

# 3. Run the Go App locally (Make sure dev-db is running first!)
start: dev-db
	@go run main.go serve

# ==============================================================================
# Testing
# ==============================================================================

# Run standard unit tests
test:
	go test ./... -v

# Run standard unit tests
test:
	go test ./... -v

# Run integration tests (Requires dev-db to be running)
test-integration:
	TEST_INTEGRATION=true go test ./... -v

# ==============================================================================
# Production Workflow
# Strategy: Run Everything (App + DBs) inside Docker
# ==============================================================================

# Build the Docker image and start everything in production mode
prod:
	docker compose -f docker-compose.prod.yml up -d --build

# Stop the production containers
prod-stop:
	docker compose -f docker-compose.prod.yml down

# Check logs in production
prod-logs:
	docker compose -f docker-compose.prod.yml logs -f

# ==============================================================================
# Utility
# ==============================================================================

# Compile the binary locally (good for checking compile errors without running)
build:
	CGO_ENABLED=0 go build -o bin/bot main.go

clean: dev-db-stop
	rm -rf bin/