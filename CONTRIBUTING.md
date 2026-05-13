# Contributing to the Inference Gateway Go SDK

Thank you for your interest in contributing to the [Inference Gateway Go SDK](sdk.go)! This document provides guidelines and steps for contributing.

## Table of Contents

- [Contributing to the Inference Gateway Go SDK](#contributing-to-the-inference-gateway-go-sdk)
  - [Table of Contents](#table-of-contents)
  - [Development Setup](#development-setup)
  - [Development Process](#development-process)
  - [Pull Request Process](#pull-request-process)
  - [Release Process](#release-process)
  - [Getting Help](#getting-help)

## Development Setup

The project ships a [Flox](https://flox.dev) environment that pins Go and `golangci-lint` to the versions used in CI. Using Flox is the recommended path; a manual setup is also documented below.

### Option A — Flox (recommended)

1. Install [Flox](https://flox.dev/docs/install-flox/) and [Task](https://taskfile.dev).
2. Clone and activate the environment:

```sh
git clone https://github.com/inference-gateway/sdk
cd sdk
flox activate
go mod download
```

`flox activate` provides the pinned Go toolchain and `golangci-lint` for the current shell. Re-run it whenever you enter the project in a new shell.

### Option B — Manual setup

1. Install:

-   [Go](https://go.dev/dl/) 1.26 or newer
-   [Task](https://taskfile.dev)
-   [golangci-lint](https://golangci-lint.run)
-   Git

2. Clone the repository and download dependencies:

```sh
git clone https://github.com/inference-gateway/sdk
cd sdk
go mod download
```

### Regenerating types (optional)

If the upstream OpenAPI spec has changed:

```sh
task oas-download
task generate
```

## Development Process

1. Create a new branch:

```sh
git checkout -b my-feature
```

2. Make changes and run the tests:

```sh
task test
```

3. Add tests for new features or fix tests for refactoring and bug fixes.

4. Run the linter:

```sh
task lint
```

5. Format your code before committing:

```sh
go fmt ./...
```

## Pull Request Process

1. Commit changes:

```sh
git add .
git commit -m "Add my feature"
```

Types:

-   feat: new feature
-   fix: bug fix
-   refactor: code change that neither fixes a bug nor adds a feature
-   docs: documentation
-   style: formatting, missing semi colons, etc; no code change
-   test: adding missing tests
-   chore: updating build tasks, package manager configs, etc; no production code change
-   ci: changes to CI configuration files and scripts
-   perf: code change that improves performance

2. Ensure your PR:

-   Passes all tests
-   Updates documentation as needed
-   Includes tests for new features
-   Has a clear description of changes

## Release Process

1. Merging to main triggers CI checks
2. Manual release workflow can be triggered from Actions
3. Version is determined by commit messages
4. Changelog is automatically generated

## Getting Help

-   File an issue in [GitHub Issues](https://github.com/inference-gateway/sdk/issues)
-   For questions about the API, consult the [openapi.yaml](openapi.yaml) specification
