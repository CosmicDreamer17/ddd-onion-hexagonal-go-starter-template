# AI Agent Protocols (Claude)

## Core Architecture
- **Monorepo:** Go Backend (`/backend`) + Next.js Frontend (`/frontend`).
- **Backend Pattern:** Hexagonal (Onion) with Layered Isolation.
  - `internal/domain`: Pure business logic, no external dependencies.
  - `internal/application`: Use cases, interfaces (Ports).
  - `internal/infra`: Implementations (Adapters), Database, External APIs.
  - `internal/api`: HTTP Handlers, DTOs.
- **Database:** SQLite (`mattn/go-sqlite3`).
- **Generation:**
  - `sqlc`: Use `query.sql` for all DB interactions. Never write raw SQL strings in Go.
  - `tygo`: Generates TypeScript interfaces from `backend/internal/api/dtos.go`.

## Mandatory Rules
- **Layer Isolation:** `domain` layer MUST NOT import from `infra` or `api`. Verified via `golangci-lint` (depguard).
- **Plan First:** Always use `/plan` mode before modifying the Domain layer.
- **Type Safety:** Never use raw strings for domain IDs; always use defined custom types (e.g., `domain.UserID`).
- **Error Context:** Errors must be wrapped with context: `fmt.Errorf("failed to do X: %w", err)`.
- **Database Safety:** All SQL must reside in `query.sql`. No raw SQL strings in Go.
- **Verification:** Architecture is verified via `scripts/verify.sh`. This script is the source of truth.

## Commands
- `make generate`: Synchronizes DB and TypeScript types.
- `make verify`: Full CI check locally.
- `make dev`: Start development environment.
