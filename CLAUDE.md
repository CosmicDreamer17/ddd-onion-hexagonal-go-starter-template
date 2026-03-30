# Claude Code Onboarding

**CRITICAL**: Read [AI.md](./AI.md) for the full architectural mandates and identity rules before modifying this repository.

## 🛠 Project Context
- **Primary Language**: Go (Backend), TypeScript (Frontend)
- **Architecture**: Hexagonal + DDD Monorepo
- **Verification**: `make verify`

## 📡 Key Commands
- `make generate`: Synchronizes DB and TypeScript types.
- `make verify`: Full CI check locally.
- `make dev`: Start development environment.
