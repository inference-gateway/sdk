# Repository Guidelines

## Project Structure & Module Organization

This repository contains the Go SDK for Inference Gateway. Core package code lives at the repository root:

- `sdk.go` contains the client interface and implementation; `types.go` contains SDK helper types.
- `generated_types.go` is generated from `openapi.yaml`; avoid manual edits unless regenerating.
- `sdk_test.go` and `vision_test.go` contain root package tests.
- `examples/` contains standalone Go modules for generation, streaming, tools, vision, models, reasoning, MCP tool listing, and middleware bypass examples.

OpenAPI generation is configured by `oapi-codegen.yaml`. Automation is defined in `Taskfile.yml`, and CI runs lint, build, and tests.

## Build, Test, and Development Commands

Use Flox when available to get the pinned Go and lint tooling:

```sh
flox activate
go mod download
```

Common commands:

- `task test` runs `go test -v ./...`.
- `task lint` runs `golangci-lint run`.
- `go build -v ./...` builds all packages, matching CI.
- `task tidy` runs `go mod tidy` in each module, including examples.
- `task oas-download` refreshes `openapi.yaml` from the schemas repository.
- `task generate` regenerates `generated_types.go` and applies `gofmt`.
- `task docs` starts local Go documentation on `:6060`.

## Coding Style & Naming Conventions

Follow Go conventions and run `gofmt` before committing. `.editorconfig` requires tabs for Go files and two-space indentation for YAML, JSON, TOML, and most other files. Keep public identifiers clear and documented when exported. Tests use Go's `TestName` naming pattern and `testify/assert` or `testify/require`.

Do not hand-edit generated OpenAPI types when the schema changed; update `openapi.yaml` and run `task generate`.

## Testing Guidelines

Add or update tests for new client behavior, request construction, response parsing, retries, streaming, and error handling. Prefer `httptest` servers for HTTP behavior so tests remain deterministic. Run `task test` locally before opening a PR; run `task lint` when touching Go code.

## Commit & Pull Request Guidelines

The project uses Conventional Commit-style messages, such as `feat:`, `fix:`, `docs:`, `test:`, `chore:`, and scoped variants like `chore(docs):`. Release notes and versions are derived from commit messages, so keep them specific and accurate.

Pull requests should include a concise description, linked issues when relevant, tests for behavior changes, and documentation updates for user-facing changes. Confirm that lint, build, and tests pass before requesting review.

## Security & Configuration Tips

Do not commit real API keys or local secrets. Keep environment-specific values in `.env` or shell configuration. Use placeholder tokens in examples and document required variables in the example README.
