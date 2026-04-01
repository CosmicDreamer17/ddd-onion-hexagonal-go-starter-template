#!/bin/bash
set -e

echo "Running architectural verification..."

# 1. Backend checks
echo "==> Backend"
cd backend
go mod tidy
golangci-lint run
govulncheck ./...
CGO_ENABLED=1 go test -race -count=1 ./...

# 2. Frontend checks
echo "==> Frontend"
cd ../frontend
npm ci --prefer-offline
npm run lint
npm run build

echo "All checks passed!"
