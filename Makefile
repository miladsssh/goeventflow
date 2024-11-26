# Makefile for GoEventFlow

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test
GOBUILD = $(GOCMD) build
GOFMT = $(GOCMD) fmt
GOVET = $(GOCMD) vet
GOLINT = golangci-lint

# Directories
BINDIR = bin

# Binary name
BINARY_NAME = goeventflow

# Default target
all: fmt vet lint test build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOCMD) mod download

# Format the code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Vet the code
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Lint the code
lint:
	@echo "Linting code..."
	$(GOLINT) run ./...

# Build the application
build:
	@echo "Building application..."
	$(GOBUILD) -o $(BINDIR)/$(BINARY_NAME) ./examples/main.go

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BINDIR)

# Run the application
run: build
	@echo "Running application..."
	@./$(BINDIR)/$(BINARY_NAME)

# Phony targets
.PHONY: all deps fmt vet lint build test clean run help
