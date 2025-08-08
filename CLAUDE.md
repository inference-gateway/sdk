# CLAUDE.md

## Commands
- `task oas-download` - Download OpenAPI spec
- `task generate` - Generate types from OpenAPI
- `task test` - Run tests  
- `task lint` - Lint code
- `task docs` - Generate docs

## Development
- SDK based on `openapi.yaml` spec
- Follow TDD: write tests first, then implement
- Use table-driven tests for related scenarios
- All changes via Pull Requests

## Code Style
- Simple, clear Go code with standard conventions
- Use context for cancellation/timeouts
- Explicit error handling with descriptive messages
- Format with `go fmt ./...` before committing
