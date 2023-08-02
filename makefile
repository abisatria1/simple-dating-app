# Variables
APP_NAME = simple-bumble
MAIN_FILE = cmd/api/application.go
MIGRATION_FILE = cmd/migration/migration.go

.DEFAULT_GOAL := run

# Install project dependencies
install:
	go mod download

# Build the application
build:
	go build -o $(APP_NAME) $(MAIN_FILE)

# Run the application using 'air'
run:
	air

# Run migration 
migration: 
	go run ${MIGRATION_FILE}

# Clean up the build artifacts
clean:
	rm -f $(APP_NAME)

# Run tests
test:
	go test ./... -race -count=1

.PHONY: install build run clean test
