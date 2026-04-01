# Codex / OpenAI Onboarding

**CRITICAL**: Read [AI.md](./AI.md) for the full architectural mandates and identity rules before modifying this repository.

## 🛠 Project Context
- **Architecture**: Hexagonal + DDD Monorepo
- **Primary Stack**: Go (Backend), Next.js (Frontend)
- **Verification**: Run `make verify` to ensure shippability.

## 📡 Key Commands
- `make generate`: Synchronizes DB and TypeScript types.
- `make verify`: Full CI check locally.
- `make dev`: Start development environment.
