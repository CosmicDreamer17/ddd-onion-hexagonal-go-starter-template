.PHONY: generate verify dev

TYGO := ~/go/bin/tygo

generate:
	cd backend && sqlc generate
	cd backend && $(TYGO) generate

verify:
	cd backend && go mod tidy
	cd backend && golangci-lint run
	cd backend && govulncheck ./...
	cd backend && go test ./...

dev:
	# Concurrent dev server (simplified)
	(cd backend && go run cmd/api/main.go) & (cd frontend && npm run dev)
