# AGENTS.md

## Project Structure & Module Organization

The package is intentionally flat at the repository root.

- `generated_types.go` is generated from `openapi.yaml` - never hand-edit it (marked `linguist-generated` in `.gitattributes`). `openapi.yaml` is a local copy of the upstream schema - do not edit it either. To change types, change the upstream schema in `inference-gateway/schemas`, then run `task oas-download && task generate` (pin the schema with `SCHEMAS_REF`).
- `examples/` contains standalone Go modules (each with its own `go.mod`/`go.sum`). They are not built by `go build ./...` from the root. If SDK changes affect them, update each example and run `task tidy`.

## Build, Test, and Development Commands

The toolchain (Go, golangci-lint, Task) is pinned in `.flox/env/manifest.toml` - use Flox when available (`flox activate`). Tasks are defined in `Taskfile.yml`. Run `task test` locally before opening a PR. Run `task lint` when touching Go code. `task precommit:install` points git at `.githooks/` - the pre-commit hook blocks commits whose staged text files lack a final newline (it appends the newline, then you `git add` and commit again). `.flox/` is skipped.

## Key Behaviors and Gotchas

- **`WithX` methods mutate the receiver.** `WithAuthToken`, `WithTools`, `WithOptions`, `WithHeaders`, `WithHeader`, and `WithMiddlewareOptions` return `Client` for chaining but do not clone, so they affect every later call on that client. For per-call config, chain inline: `client.WithTools(t).GenerateContent(...)`. `WithMiddlewareOptions` both sets headers (flag true) and deletes them (flag false).
- **Retry layer.** Every API call goes through `executeWithRetry`. Defaults: 3 attempts, 2s initial backoff, 30s cap, multiplier 2. It retries 408/429/500/502/503/504 and network errors classified by `isRetryableError`. On 429 it honors `Retry-After` (seconds or HTTP-date). Disable with `RetryConfig{Enabled: false}` or `MaxAttempts` below 1. The real fields are on `RetryConfig` (`Enabled`, `MaxAttempts`, `InitialBackoffSec`, `MaxBackoffSec`, `BackoffMultiplier`, `RetryableStatusCodes`, `OnRetry`). Any `RetryOptions`/`MaxRetries`/`MinDelay` naming in docs is stale.
- **Streaming.** `GenerateContentStream` returns a buffered (100) `<-chan SSEvent`. Callers range over it and switch on `*event.Event`: `ContentDelta` (raw JSON chunk in `event.Data`, a `*[]byte`), `StreamEnd` (emitted after `data: [DONE]`), and a nil `Event` whose `event.Data` holds `{"error": "..."}` on a read error. The channel closes when the stream ends or on read error. The `SSEventEvent` values are content-start, content-delta, content-end, message-start, message-end, stream-start and stream-end.
- **Middleware bypass.** `MiddlewareOptions.SkipMCP` maps to `X-MCP-Bypass: true` and `DirectProvider` to `X-Direct-Provider: true`. Setting the raw headers via `WithHeader` is equivalent.

## Coding Style & Naming Conventions

Every non-standard-library import is aliased to its last path element (`sdk "github.com/inference-gateway/sdk"`). Import order and aliases are enforced by `gci` and `importas` in `.golangci.yml`. Pin an alias there only when two packages would collide. Fix locally with `golangci-lint fmt` and `golangci-lint run --fix`.

### Code Readability

- Write self-explanatory code: clear names and small, single-purpose functions carry the intent.
  If a block needs a comment to be understood, extract it into a well-named function or variable.
- No inline comments inside function bodies.
- Doc comments on functions and types are at most 5 lines: what it does and why, not how.
- No comments above modules, packages, or files.
- Tool directives are not comments and stay where the tool needs them (lint suppressions, build
  tags, compiler pragmas, code generation markers).
- No semicolons in documentation prose (Markdown files, doc comments): split the sentence or use
  a dash instead.

## Commit & Pull Request Guidelines

Conventional Commits are required - release tooling (`.releaserc.yaml`, `semantic-release`) parses them. `feat` is a minor bump, while `fix`, `refactor`, `perf`, `ci`, `docs`, `style`, `test`, `build`, and `chore` are patch bumps. `chore(release): ...` is reserved for the release bot. Keep messages specific and accurate.

Releases run via the manual `Release` workflow. Never bump versions by hand or edit `CHANGELOG.md` - it is generated.

Documentation must never reference GitHub issues or tickets - do not include issue numbers, issue URLs, or cross-reference lines such as "Closes" followed by an issue. Describe the behavior, configuration, or rationale directly instead.

## Security & Configuration Tips

Do not commit real API keys or local secrets. Keep environment-specific values in `.env` or shell configuration. Use placeholder tokens in examples and document required variables in the example README.
