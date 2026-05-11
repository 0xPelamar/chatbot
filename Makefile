start:
	@go run main.go serve

# Run standard unit tests
test:
	go test ./... -v

# Run integration tests (Requires dev-db to be running)
test-integration:
	INTEGRATION_TEST=true go test ./... -v