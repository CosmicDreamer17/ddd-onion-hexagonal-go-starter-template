# AI-Autonomous Go + Next.js Monorepo Starter

A high-performance, strictly bounded Monorepo template optimized for AI-agent maintenance.

## 🚀 Quick Start (Create Your Own)

### Path A: Using GitHub CLI (Recommended)
Bootstrap a new repo from this template:
```bash
gh repo create my-new-app --template CosmicDreamer17/ddd-onion-hexagonal-go-starter-template --public --clone
```

### Path B: Using `degit` (No GitHub CLI)
```bash
npx degit CosmicDreamer17/ddd-onion-hexagonal-go-starter-template my-new-app
```

> **🤖 Note for AI Agents**: If you are bootstrapping this project, refer to the [Agentic Bootstrap](./AI.md#🚀-agentic-bootstrap-creating-a-new-project) section in `AI.md` for mandatory initialization steps.

---

## 🏗 Architecture (Hexagonal + DDD)

### Backend (Go in `/backend`)
- **`internal/domain`**: Pure business logic. Zero internal dependencies.
- **`internal/application`**: Ports (Interfaces) and Use Cases.
- **`internal/infra`**: SQLite + sqlc adapters.
- **`internal/api`**: Chi-based HTTP server + DTOs.
- **`cmd/admin`**: CLI Admin tool (Terminal Adapter).

### Frontend (Next.js App in `/frontend`)
- **Types**: Zero-drift bindings generated via `tygo` into `/frontend/src/types/generated/`.

## 📡 Operations

### 1. Development (Local)
```bash
make dev
```
- **Frontend (Next.js)**: [http://localhost:3000](http://localhost:3000)
- **Backend (Go)**: [http://localhost:8080](http://localhost:8080)

### 2. Admin CLI (Direct Use-Case Execution)
Interact with the backend directly from the terminal without HTTP overhead:
```bash
make admin ARGS="register --email user@example.com --password securepass"
```

### 3. Verification (CI/CD)
```bash
make verify
```
Runs architectural leak checks (depguard), formatting, and workspace tests.

### 4. Sync Types
```bash
make generate
```
Triggers `sqlc` and `tygo` generation.

## 🤖 Agentic Maintenance
This repository is "AI-Autonomous Ready." It includes specialized markdown files in the root:
- **`AI.md` (Master)**: The single source of truth for architectural mandates and quality-of-life instructions for all AI agents (Gemini, Claude, Codex/Cursor).
- **Tool Entry Points**: `GEMINI.md`, `CLAUDE.md`, and `CODEX.md` provide tool-specific onboarding while referencing the master rules.

## 🛠 Provenance
This repository was **autonomously engineered and verified** by **Gemini CLI** (Gemini 2.0 Flash) on **Sunday, March 29, 2026**. See [PROVENANCE.md](./PROVENANCE.md) for full metadata.
