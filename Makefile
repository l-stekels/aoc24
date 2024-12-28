.PHONY: help build clean test fmt vet lint run tidy vendor update

# Variables
BINARY_NAME := advent2024

help:
	@echo Usage: make [target]
	@powershell -Command "Get-Content Makefile | ForEach-Object { if ($$_ -match '^[a-zA-Z_-]+:.*?## .*$$') { $$target = $$_ -replace '^([a-zA-Z_-]+):.*?## (.*)$$', '  $$1 : $$2'; Write-Host $$target } }"

build: ## Compile the Go project into a binary
	go build -o $(BINARY_NAME) ./cmd/main.go

clean: ## Remove compiled binary and clean up build artifacts
	go clean
	@if exist $(BINARY_NAME) del $(BINARY_NAME)

test: ## Run all tests in the project
	go test -v ./...

fmt: ## Format all Go files in the project
	go fmt ./...

vet: ## Run go vet to check for potential issues
	go vet ./...

run: ## Run the Go application
	go run $(GOFILES)

tidy: ## Add missing and remove unused modules
	go mod tidy

vendor: ## Create a vendor directory and vendor all dependencies
	go mod vendor

update: ## Update all dependencies to their latest versions
	go get -u ./...
	go mod tidy
