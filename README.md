# Autonomous Architecture Template

A high-performance, locally-verifiable, and rigidly isolated template designed for AI-driven software engineering. 

**Stack:** Go (Hexagonal) + Next.js (App Router) + SQLite + Tailwind CSS.

## 🏗 Architecture

### Strict Hexagonal Boundaries (Backend)
Enforced locally and in CI via `depguard` and `golangci-lint`:
- `domain`: Pure Go business rules. **Zero dependencies**.
- `application`: Use-cases & interface definitions (Ports).
- `infra`: Technical implementations (Adapters - DB, external APIs).
- `api`: HTTP delivery layer (Controllers, JSON decoding).

### The "Zero-Drift" Type Bridge
- **Database to Go:** `sqlc` transpiles raw SQL (`query.sql`) into type-safe Go data access functions. 
- **Go to TypeScript:** `tygo` converts Go DTOs (`backend/internal/api/dtos`) directly into TypeScript interfaces (`frontend/src/types/generated/index.ts`).

## 🚀 Quickstart

### 1. Requirements
- [Go](https://go.dev/dl/) 1.22+
- [Node.js](https://nodejs.org/) 20+
- [Docker](https://www.docker.com/) (Optional)

### 2. Local Development

Run the concurrent dev server (requires `make`):

```bash
make dev
```
- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`

### 3. Verification & Generation

Whenever you modify SQL schemas or Go DTOs, regenerate the type bridges:

```bash
make generate
```

Before committing, run the full architectural and linting suite:

```bash
make verify
```

### 4. Docker Deployment

```bash
docker-compose up --build -d
```

## 🤖 Agent Protocols

This repository contains strict instructions for autonomous agents (like Gemini or Claude) located in `GEMINI.md`, `CLAUDE.md`, and `.gemini/skills/architect.md`. 
Agents will automatically read these instructions and act as **Principal Systems Architects**, refusing to break boundary layers and heavily utilizing `/plan` mode before touching the Domain.
