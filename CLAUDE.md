# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

The Inference Gateway Go SDK — a thin Go client (`github.com/inference-gateway/sdk`) for the Inference Gateway HTTP API, a unified front for multiple LLM providers (Ollama, OllamaCloud, Groq, OpenAI, DeepSeek, Cloudflare, Cohere, Anthropic, Google, Mistral). The SDK wraps chat completions (sync + SSE streaming), model/MCP-tool listing, vision (multimodal) messages, function/tool calling, middleware bypass, and a built-in retry layer.

## Commands

The project uses [Task](https://taskfile.dev) (see `Taskfile.yml`). Toolchain is pinned via Flox in `.flox/env/manifest.toml` (Go ^1.26.4, golangci-lint ^2.12.2) — `flox activate` from the repo root puts the right versions on `PATH`. CI runs `golangci-lint run --timeout 5m`, `go build -v ./...`, and `go test -v ./...`.

- `task test` — run the full Go test suite (`go test -v ./...`).
- Run a single test: `go test -v -run TestListModels ./...` (or `TestStream`, `TestVision`, etc.).
- `task lint` — `golangci-lint run`.
- `task tidy` — `go mod tidy` recursively (root + every example, each example has its own module).
- `task oas-download` — fetch `openapi.yaml` from `github.com/inference-gateway/schemas` (main branch).
- `task generate` — regenerate `generated_types.go` from `openapi.yaml` via `oapi-codegen` (then rewrites `interface{}` → `any`). Run after `task oas-download` when the upstream schema changes.
- `task docs` — serve godoc at `http://localhost:6060/pkg/github.com/inference-gateway/sdk`.
- Running an example: `cd examples/<name> && go run main.go` (each example is its own module; set `INFERENCE_GATEWAY_URL`, optionally `LLM_PROVIDER`, `LLM_MODEL`, and API keys for the gateway).

## Architecture

The package is intentionally flat — three hand-written files plus one generated one.

- **`sdk.go`** — the entire client. Defines the `Client` interface, the `clientImpl` struct, `NewClient`, all `WithX` builder methods, `executeWithRetry` (the retry/backoff engine), and the public API: `ListModels`, `ListProviderModels`, `ListTools`, `GenerateContent`, `GenerateContentStream`, `HealthCheck`. SSE parsing for streaming lives at the bottom of `GenerateContentStream` — it reads `data: ` lines off the raw body, emits `ContentDelta` events, and closes on `[DONE]`.
- **`types.go`** — hand-written types that aren't in the OpenAPI: `ClientOptions`, `RetryConfig`, `MiddlewareOptions`, plus the message/content-part constructor helpers (`NewMessageContent`, `NewTextMessage`, `NewImageMessage`, `NewTextContentPart`, `NewImageContentPart`). The helpers wrap the oapi-codegen-generated `oneOf` machinery so callers don't deal with `FromMessageContent0/1` directly.
- **`generated_types.go`** — **generated, do not hand-edit.** Marked `linguist-generated=true` in `.gitattributes`. Contains every request/response struct (`CreateChatCompletionRequest`, `CreateChatCompletionResponse`, `Message`, `MessageContent` union, `ContentPart`, `ChatCompletionTool`, `FunctionObject`, `Model`, `MCPTool`, `Error`, `SSEvent`, …) and the enums (`Provider`, `MessageRole`, `FinishReason`, `ImageURLDetail`, `SSEventEvent`, …) with `Valid()` methods. To change anything in here, edit the upstream OpenAPI in `inference-gateway/schemas`, then `task oas-download && task generate`.
- **`openapi.yaml`** — local copy of the upstream schema. Don't edit; refresh via `task oas-download`.
- **`oapi-codegen.yaml`** — generator config: package `sdk`, output `generated_types.go`, `ToCamelCaseWithInitialisms` name normalizer.

Two test files: `sdk_test.go` (~33 tests, covers client/list/generate/stream/retry/middleware paths against an `httptest.NewServer` mock) and `vision_test.go` (multimodal/content-part-specific cases). There is no separate integration suite — everything is mock-based.

### Builder pattern — `WithX` methods MUTATE

`WithAuthToken`, `WithTools`, `WithOptions`, `WithHeaders`, `WithHeader`, `WithMiddlewareOptions` all return `Client` for chaining but modify the receiver. They do not clone. `client.WithAuthToken("x")` changes every subsequent call on `client` (and on anything aliasing it). When a user wants per-call config, they generally chain inline: `client.WithTools(t).GenerateContent(...)`. `WithMiddlewareOptions` is special — it both sets headers (when the flag is true) and *deletes* them (when false), so it can be used to clear prior middleware state.

### Retry layer

Lives entirely in `executeWithRetry` in `sdk.go`. Every API call goes through it. Defaults: 3 attempts, 2s initial backoff, 30s cap, multiplier 2, retries 408/429/500/502/503/504 and network errors classified by `isRetryableError` (timeouts, ECONNREFUSED, ECONNRESET, DNS errors, EOF). On 429 it parses `Retry-After` (seconds or HTTP-date) and waits that long instead of the computed backoff. Disable with `RetryConfig{Enabled: false}` or by setting `MaxAttempts: 0` via `RetryConfig`. (The README's `RetryOptions` / `MaxRetries` / `MinDelay` naming is stale — the real fields are on `RetryConfig`: `Enabled`, `MaxAttempts`, `InitialBackoffSec`, `MaxBackoffSec`, `BackoffMultiplier`, `RetryableStatusCodes`, `OnRetry`.)

### Streaming

`GenerateContentStream` returns `<-chan SSEvent` (buffered 100) and a goroutine that reads the raw HTTP body. Callers `for event := range events` and switch on `*event.Event` between `ContentDelta` (payload in `event.Data`, a `*[]byte` of the raw JSON SSE chunk), `StreamEnd` (after `data: [DONE]`), and `MessageError`. The channel is closed when the stream ends or on read error.

### Middleware bypass

Two boolean flags on `MiddlewareOptions` map to HTTP headers the gateway recognises: `SkipMCP` → `X-MCP-Bypass: true`, `DirectProvider` → `X-Direct-Provider: true`. Either set them via `WithMiddlewareOptions` or set the raw headers via `WithHeader` — equivalent behaviour.

### Examples

Each subdirectory under `examples/` is its own Go module (its own `go.mod`/`go.sum`) so it can be `go run` standalone. They are not part of the root module's build and aren't covered by `go build ./...` from the root. If you change SDK types in a way that affects them, update each example separately and run `task tidy` to refresh all `go.mod` files.

## Conventions

- **Conventional commits** are required — release tooling (`.releaserc.yaml`) parses them. Recognised types: `feat` (minor bump), `fix` / `refactor` / `perf` / `impr` / `ci` / `docs` / `style` / `test` / `build` / `chore` (patch). `impr` is a project-specific "improvement" type — use it for small enhancements that aren't quite features. `chore(release): ...` is reserved for the release bot.
- Releases run via the manual `Release` workflow (`workflow_dispatch`), driven by `semantic-release`. Don't bump versions by hand or write to `CHANGELOG.md` — it's generated.
- Go formatting: tabs (per `.editorconfig`); always `go fmt ./...` before committing.
- Never hand-edit `generated_types.go` or `CHANGELOG.md`.
