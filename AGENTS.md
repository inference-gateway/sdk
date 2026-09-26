# AGENTS.md

## Project Structure & Module Organization

This repository contains the Go SDK (`github.com/inference-gateway/sdk`) for the Inference Gateway, a unified HTTP API in front of many LLM providers. The package is intentionally flat at the repository root:

- `sdk.go` contains the `Client` interface, `clientImpl`, `NewClient`, all `WithX` builder methods, the retry engine (`executeWithRetry`), and every public API call, including SSE stream parsing.
- `types.go` contains hand-written types not in the OpenAPI (`ClientOptions`, `RetryConfig`, `MiddlewareOptions`) and constructor helpers (`NewTextMessage`, `NewImageContentPart`, ...) that hide the generated `oneOf` machinery.
- `generated_types.go` is generated from `openapi.yaml` - never hand-edit it (marked `linguist-generated` in `.gitattributes`). `openapi.yaml` is a local copy of the upstream schema - do not edit it either; refresh it via `task oas-download`.
- `*_test.go` files are all mock-based (`httptest.NewServer`); there is no separate integration suite.
- `examples/` contains standalone Go modules (each with its own `go.mod`/`go.sum`). They are not built by `go build ./...` from the root; if SDK changes affect them, update each example and run `task tidy`. Run one with `cd examples/<name> && go run main.go` (set `INFERENCE_GATEWAY_URL`, optionally `LLM_PROVIDER`, `LLM_MODEL`, and gateway API keys).

OpenAPI generation is configured by `oapi-codegen.yaml`. Automation is defined in `Taskfile.yml`, and CI runs `golangci-lint run --timeout 5m`, `go build -v ./...`, and `go test -v ./...`.

## Build, Test, and Development Commands

The toolchain (Go, golangci-lint, Task) is pinned in `.flox/env/manifest.toml`; use Flox when available:

```sh
flox activate
go mod download
```

Common commands:

- `task test` runs `go test -v ./...`. Run a single test with `go test -v -run TestListModels ./...`.
- `task lint` runs `golangci-lint run`.
- `go build -v ./...` builds all packages, matching CI.
- `task build-examples` builds every standalone example module.
- `task tidy` runs `go mod tidy` in each module, including examples.
- `task oas-download` refreshes `openapi.yaml` from the schemas repository (pin with `SCHEMAS_REF`).
- `task generate` regenerates `generated_types.go` and rewrites `interface{}` to `any`. `task oas-sync` does both.
- `task docs` starts local Go documentation on `:6060`.

## Key Behaviors and Gotchas

- **`WithX` methods mutate the receiver.** `WithAuthToken`, `WithTools`, `WithOptions`, `WithHeaders`, `WithHeader`, and `WithMiddlewareOptions` return `Client` for chaining but do not clone, so they affect every later call on that client. For per-call config, chain inline: `client.WithTools(t).GenerateContent(...)`. `WithMiddlewareOptions` both sets headers (flag true) and deletes them (flag false).
- **Retry layer.** Every API call goes through `executeWithRetry`. Defaults: 3 attempts, 2s initial backoff, 30s cap, multiplier 2; retries 408/429/500/502/503/504 and network errors classified by `isRetryableError`. On 429 it honors `Retry-After` (seconds or HTTP-date). Disable with `RetryConfig{Enabled: false}` or `MaxAttempts` below 1. The real fields are on `RetryConfig` (`Enabled`, `MaxAttempts`, `InitialBackoffSec`, `MaxBackoffSec`, `BackoffMultiplier`, `RetryableStatusCodes`, `OnRetry`); any `RetryOptions`/`MaxRetries`/`MinDelay` naming in docs is stale.
- **Streaming.** `GenerateContentStream` returns a buffered (100) `<-chan SSEvent`. Callers range over it and switch on `*event.Event`: `ContentDelta` (raw JSON chunk in `event.Data`, a `*[]byte`), `StreamEnd` (after `data: [DONE]`), and `MessageError`. The channel closes when the stream ends or on read error.
- **Middleware bypass.** `MiddlewareOptions.SkipMCP` maps to `X-MCP-Bypass: true` and `DirectProvider` to `X-Direct-Provider: true`; setting the raw headers via `WithHeader` is equivalent.

## Coding Style & Naming Conventions

Follow Go conventions and run `gofmt` before committing. `.editorconfig` requires tabs for Go files and two-space indentation for YAML, JSON, TOML, and most other files. Keep public identifiers clear and documented when exported. Tests use Go's `TestName` naming pattern and `testify/assert` or `testify/require`.

Import order is enforced by the `gci` formatter (see `.golangci.yml`): standard library, `github.com/stretchr/testify`, third-party, `github.com/inference-gateway/*`, then this module. Every non-standard-library import must be named after its last path element (`sdk "github.com/inference-gateway/sdk"`), enforced by `importas`; pin an alias in `.golangci.yml` only when two packages would collide. Fix locally with `golangci-lint fmt` and `golangci-lint run --fix`.

Do not hand-edit generated OpenAPI types; change the upstream schema in `inference-gateway/schemas`, then run `task oas-download && task generate`.

### Code Readability

- Write self-explanatory code: clear names and small, single-purpose functions carry the intent.
  If a block needs a comment to be understood, extract it into a well-named function or variable.
- No inline comments inside function bodies.
- Doc comments on functions and types are at most 5 lines: what it does and why, not how.
- No comments above modules, packages, or files.
- Tool directives are not comments and stay where the tool needs them (lint suppressions, build
  tags, compiler pragmas, code generation markers).

## Testing Guidelines

Add or update tests for new client behavior, request construction, response parsing, retries, streaming, and error handling. Prefer `httptest` servers for HTTP behavior so tests remain deterministic. Run `task test` locally before opening a PR; run `task lint` when touching Go code.

## Commit & Pull Request Guidelines

Conventional Commits are required - release tooling (`.releaserc.yaml`, `semantic-release`) parses them. `feat` is a minor bump; `fix`, `refactor`, `perf`, `impr`, `ci`, `docs`, `style`, `test`, `build`, and `chore` are patch bumps. `impr` is a project-specific type for small enhancements that are not quite features. `chore(release): ...` is reserved for the release bot. Keep messages specific and accurate.

Releases run via the manual `Release` workflow. Never bump versions by hand or edit `CHANGELOG.md` - it is generated.

Pull requests should include a concise description, tests for behavior changes, and documentation updates for user-facing changes. Confirm that lint, build, and tests pass before requesting review.

Documentation must never reference GitHub issues or tickets - do not include issue numbers, issue URLs, or cross-reference lines such as "Closes" followed by an issue. Describe the behavior, configuration, or rationale directly instead.

## Security & Configuration Tips

Do not commit real API keys or local secrets. Keep environment-specific values in `.env` or shell configuration. Use placeholder tokens in examples and document required variables in the example README.
