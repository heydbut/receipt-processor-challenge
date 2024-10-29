APP_NAME := receipt-processor
PKG := ./...
BIN_DIR := bin
BIN_NAME := $(BIN_DIR)/$(APP_NAME)
COVERAGE_OUT := coverage.out
COVERAGE_HTML := coverage.html

all: test build

test:
	go test -v $(PKG)

run:
	go run cmd/main.go

coverage:
	go test -coverprofile=$(COVERAGE_OUT) $(PKG)
	go tool cover -html=$(COVERAGE_OUT) -o $(COVERAGE_HTML)
	@echo "Code coverage report generated at $(COVERAGE_HTML)"

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_NAME) cmd/main.go
	@echo "Application built at $(BIN_NAME)"

docker-build:
	docker build -t $(APP_NAME) .

lint:
	golangci-lint run

format:
	go fmt $(PKG)

clean:
	rm -rf $(BIN_DIR) $(COVERAGE_OUT) $(COVERAGE_HTML)

help:
	@echo "Makefile commands:"
	@echo "  make all          - Runs tests and builds the application"
	@echo "  make test         - Runs unit tests"
	@echo "  make run          - Runs the application"
	@echo "  make coverage     - Generates code coverage report"
	@echo "  make build        - Builds the application binary"
	@echo "  make docker-build - Builds the Docker image"
	@echo "  make lint         - Runs code linter"
	@echo "  make format       - Formats the code"
	@echo "  make clean        - Cleans build artifacts"

.PHONY: all test run coverage build docker-build lint format clean help
