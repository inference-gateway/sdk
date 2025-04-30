# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Process

-   This SDK is based on the OpenAPI specification in `openapi.yaml`
-   For new features, first check `openapi.yaml` to understand the API design
-   The `openapi.yaml` changes should always be pulled from remote using `task oas-download`
-   Follow Documentation-Driven Development (DDD) and Test-Driven Development (TDD):
    1. Write documentation comments first
    2. Write tests that validate expected behavior (prefer table-driven tests for related scenarios)
    3. Implement the functionality
-   All changes should be submitted via Pull Requests following the commit conventions

## Build/Test/Lint Commands

-   Download OpenAPI spec: `task oas-download`
-   Generate types from OpenAPI: `task generate`
-   Build: `go build ./...`
-   Run all tests: `go test ./...` or `task test`
-   Run a single test: `go test -v -run TestName` (e.g., `go test -v -run TestGenerateContent`)
-   Generate documentation: `task docs` (opens on http://localhost:6060/pkg/github.com/inference-gateway/sdk)
-   Lint: `task lint` or `golangci-lint run`
-   Tidy dependencies: `task tidy` or `go mod tidy`

## Code Style Guidelines

-   Code structure should be simple, clear and easy to understand
-   Format with `go fmt ./...` before committing
-   Use explicit error handling with descriptive messages
-   Follow standard Go naming conventions (camelCase for unexported, PascalCase for exported)
-   Pointer helpers (e.g., `stringPtr`) should be used for nullable fields
-   Use context for cancellation/timeouts in all client methods
-   Include usage examples in function documentation comments
-   Error messages should be specific and descriptive
-   Use `fmt.Errorf` with context for error wrapping
-   Follow commit message conventions from CONTRIBUTING.md (feat, fix, refactor, etc.)

## API Conventions

-   Model IDs include provider prefixes (e.g., "openai/gpt-4o", "anthropic/claude-3-opus-20240229")
-   Required fields are non-pointer types (e.g., Model.Id is string not *string)
-   FunctionParameters is a map[string]interface{} with schema fields at the top level
-   Table-driven tests should be used when testing multiple related scenarios
-   All API methods should have comprehensive examples in documentation comments
