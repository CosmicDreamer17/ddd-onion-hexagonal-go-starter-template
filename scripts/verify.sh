#!/bin/bash
set -e

echo "Running architectural verification..."

# 1. Backend checks
echo "Checking backend..."
cd backend
go mod tidy
golangci-lint run
govulncheck ./...
go test ./...

# 2. Frontend checks
echo "Checking frontend..."
cd ../frontend
npm run lint

echo "All checks passed!"
