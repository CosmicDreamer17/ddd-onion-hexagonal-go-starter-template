# Skill: Principal Systems Architect (Hexagonal/Go)

This skill provides expert guidance for maintaining and extending the Hexagonal (Onion) architecture in this monorepo.

## Architectural Constraints

### 1. Domain Layer (`internal/domain`)
- **Responsibility:** Business logic and core entities.
- **Constraints:** ZERO external imports (except stdlib and approved base libs like `uuid`).
- **Guidelines:**
  - Use custom types for identifiers (e.g., `type UserID uuid.UUID`).
  - Validation logic belongs here.
  - Entities should be self-contained.

### 2. Application Layer (`internal/application`)
- **Responsibility:** Use cases and orchestration.
- **Constraints:** May import `domain`, but NOT `infra` or `api`.
- **Guidelines:**
  - Define interfaces (Ports) for persistence and external services here.
  - No database-specific logic or types.
  - Handle business transactions here.

### 3. Infra Layer (`internal/infra`)
- **Responsibility:** Technical implementations (Adapters).
- **Constraints:** Imports `application` (to implement interfaces) and `domain`.
- **Guidelines:**
  - SQL code goes into `query.sql`.
  - Maps DB models to Domain entities.
  - External API clients belong here.

### 4. API Layer (`internal/api`)
- **Responsibility:** Inbound entry points (Adapters).
- **Constraints:** Imports `application` and `domain`.
- **Guidelines:**
  - Handles HTTP, JSON, and DTOs.
  - No business logic; delegates to `application`.
  - DTOs defined in `dtos.go` are bridged to TypeScript via `tygo`.

## Workflows

### Adding a New Entity
1. Define the entity and ID type in `internal/domain`.
2. Define the Repository interface in `internal/application`.
3. Add the SQL schema and queries in `internal/infra/database`.
4. Run `make generate`.
5. Implement the interface in `internal/infra`.
6. Add the use case in `internal/application`.
7. Define DTOs and handler in `internal/api`.
8. Run `make generate` again to sync TypeScript types.
9. Verify with `make verify`.

## Enforcement
Architectural boundaries are enforced via `golangci-lint` using the `depguard` linter. Any violation will fail the build in CI.
