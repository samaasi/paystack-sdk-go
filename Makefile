.PHONY: all build test test-race fmt vet clean

# Default target: format, vet, test, and build
all: fmt vet test build

# Build all packages
build:
	go build -v ./...

# Run all tests
test:
	go test -v ./...

# Run tests with race detection (recommended for concurrency checks)
test-race:
	go test -race -v ./...

# Format code (as required by CONTRIBUTING.md)
fmt:
	go fmt ./...

# Static analysis (as required by CONTRIBUTING.md)
vet:
	go vet ./...

# Clean build artifacts
clean:
	go clean ./...