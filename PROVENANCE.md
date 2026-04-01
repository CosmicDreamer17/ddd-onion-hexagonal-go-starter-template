# Project Provenance & AI Metadata

This repository was autonomously engineered and verified by **Gemini CLI**.

## 🛠 Generation Metadata
- **Agent**: Gemini CLI (v0.35.3)
- **Primary Model**: Gemini 2.0 Flash
- **Creation Date**: Sunday, March 29, 2026
- **Architectural Pattern**: Hexagonal (Ports & Adapters) + Domain-Driven Design (DDD)
- **Verification Method**: Autonomous execution of `scripts/verify.sh` (Golangci-lint, Tests, Depguard)

## 🎯 Design Intent
The codebase was constructed using a **Research -> Strategy -> Execution** lifecycle. Every architectural boundary (e.g., Domain Purity) was programmatically verified during the build process to ensure zero technical debt at the point of delivery.

## 🤖 Machine-Readable Context
For future AI agents maintaining this project:
- This repository is a **Strict Monorepo**.
- Identity generation is **Backend-Driven** (UUID v4 in `domain`).
- Type safety is enforced via **Zero-Drift Go-to-TS bindings (tygo)**.
- Dual-entry points (Go HTTP API & Admin CLI) share the exact same Application Use Cases.
