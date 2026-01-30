.PHONY: help test test-verbose test-coverage fmt vet build clean deps tidy docs pre-commit all

# Variables
GO=go
GOTEST=$(GO) test
GOVET=$(GO) vet
GOFMT=gofmt
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

# Default target
help: ## Display this help message
	@echo "XRP X-Address Converter - Makefile Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Testing
test: ## Run all tests
	@echo "Running tests..."
	@$(GOTEST) -v ./xaddress/...

test-verbose: ## Run tests with verbose output
	@echo "Running tests with verbose output..."
	@$(GOTEST) -v -race ./xaddress/...

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@$(GOTEST) -v -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./xaddress/...
	@$(GO) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"
	@$(GO) tool cover -func=$(COVERAGE_FILE)

# Code Quality
fmt: ## Format code
	@echo "Formatting code..."
	@$(GOFMT) -s -w .

vet: ## Run go vet
	@echo "Running go vet..."
	@$(GOVET) ./...

check: fmt vet test ## Run all checks (format, vet, test)

# Build
build: ## Build the project
	@echo "Building..."
	@$(GO) build -v ./xaddress/...

# Dependencies
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@$(GO) mod download

tidy: ## Tidy and verify dependencies
	@echo "Tidying dependencies..."
	@$(GO) mod tidy
	@$(GO) mod verify

# Cleanup
clean: ## Clean build artifacts and coverage files
	@echo "Cleaning..."
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	@$(GO) clean -cache -testcache

# Documentation
docs: ## Generate and view documentation
	@echo "Opening package documentation..."
	@$(GO) doc -all ./xaddress

# Git helpers
pre-commit: fmt vet test ## Run pre-commit checks
	@echo "Pre-commit checks passed!"

# All
all: clean deps build test ## Clean, download deps, build, and test
