## [1.4.1](https://github.com/inference-gateway/sdk/compare/v1.4.0...v1.4.1) (2025-02-02)

### 📚 Documentation

* Update README with the correct examples ([6b8268d](https://github.com/inference-gateway/sdk/commit/6b8268de350580fec09d2c36a92ad131929c2518))

## [1.4.0](https://github.com/inference-gateway/sdk/compare/v1.3.0...v1.4.0) (2025-02-02)

### ✨ Features

* Add streaming content generation support to SDK ([6f4b4a0](https://github.com/inference-gateway/sdk/commit/6f4b4a046dbd011d3c7f04f0d28b6f5e96bcd73c))

### ♻️ Improvements

* Add context parameter to SDK client methods for improved context management ([cac2f4e](https://github.com/inference-gateway/sdk/commit/cac2f4edd8494b84bd7007a7a1dd67646baabdf0))
* Rename role constants for clarity in SDK message handling ([2782890](https://github.com/inference-gateway/sdk/commit/278289059f998d372c387dc85cb237b981253855))

### 📚 Documentation

* Correct typo in Anthropic provider listing in README ([0195de7](https://github.com/inference-gateway/sdk/commit/0195de724f119eacaac8412dcf3b0a5fc92822bb))
* Improve docs by having a docblocks with examples ([19f105d](https://github.com/inference-gateway/sdk/commit/19f105d528de3aa0950d48d473c222d66c0f9701))
* Remove Google LLM provider from SDK and documentation ([8ba8ec0](https://github.com/inference-gateway/sdk/commit/8ba8ec0c1b5e6c8183f74a05a3ffffb3dbd5eaf0))

### 🔧 Miscellaneous

* Cleanup ([f6547a8](https://github.com/inference-gateway/sdk/commit/f6547a8d4abb39540dcd4d3f34ef9556ad3a3d47))
* **openapi:** Update documentation for SDK usage ([f84968d](https://github.com/inference-gateway/sdk/commit/f84968d694cc3b8b805670bef020b3a417520fdc))

## [1.3.0](https://github.com/inference-gateway/sdk/compare/v1.2.3...v1.3.0) (2025-01-28)

### ✨ Features

* Add ListProviderModels to list a specific providers models ([1cf86fb](https://github.com/inference-gateway/sdk/commit/1cf86fb0debf58b54d3f5d76575fce3672c43f9e))
* Add support for Anthropic LLM provider in SDK ([fa5772a](https://github.com/inference-gateway/sdk/commit/fa5772a23e9e8d358081b48479460a24aceb2653))

### 📦 Improvements

* Add lint task to run golangci-lint in Taskfile ([60403f1](https://github.com/inference-gateway/sdk/commit/60403f1a746fb481fc04d7980025aeef91666f2c))

### 📚 Documentation

* **openapi:** Update OpenAPI spec - download the latest one from inference-gateway ([2056147](https://github.com/inference-gateway/sdk/commit/205614711a886998e7179fe4b50d16b2eacd65d7))

## [1.2.3](https://github.com/inference-gateway/sdk/compare/v1.2.2...v1.2.3) (2025-01-21)

### ♻️ Improvements

* Change Client type to interface and implement clientImpl for SDK client methods ([003b660](https://github.com/inference-gateway/sdk/commit/003b66053fdd069ddb5ee1435c211aa9b53362c2))

### 🔧 Miscellaneous

* Change semantic-release configurations build prefix to be listed under Improvements ([0ab0d02](https://github.com/inference-gateway/sdk/commit/0ab0d02c2011dfef1292ab4541ac8e5df206dd93))

## [1.2.2](https://github.com/inference-gateway/sdk/compare/v1.2.1...v1.2.2) (2025-01-21)

### 📦 Miscellaneous

* Add VSCode task extension to development container configuration for easy execution of tasks ([3592065](https://github.com/inference-gateway/sdk/commit/35920650d54122d29716d489efbc35025c13eebf))

## [1.2.1](https://github.com/inference-gateway/sdk/compare/v1.2.0...v1.2.1) (2025-01-21)

### 🐛 Bug Fixes

* Update repository references in configuration files and documentation ([39d0041](https://github.com/inference-gateway/sdk/commit/39d0041a7a2cdbac77083a40f086bf0c8a9cc5d8))

## [1.2.0](https://github.com/inference-gateway/go-sdk/compare/v1.1.0...v1.2.0) (2025-01-21)

### ✨ Features

* Add Go tools installation and enhance Taskfile with tidy, test, and docs tasks ([9c27a1d](https://github.com/inference-gateway/go-sdk/commit/9c27a1d8db4aff95da3b03a56b4ed0e911ffe1bd))

### ♻️ Improvements

* Update module name and README for improved clarity and documentation ([59ef21a](https://github.com/inference-gateway/go-sdk/commit/59ef21ae2d652707dd5b3d5a526cc8f2bcbdd355))

### 🐛 Bug Fixes

* Correct spelling of "helpful" in README and update tests to reflect changes ([ba7a423](https://github.com/inference-gateway/go-sdk/commit/ba7a4237a9f523d5e6a0421cab62eaf56e59c258))

### 📚 Documentation

* Enhance SDK documentation with detailed comments and examples for clarity ([1727352](https://github.com/inference-gateway/go-sdk/commit/17273529263efce5bac34473b8bec2fdde0d3760))

## [1.1.0](https://github.com/inference-gateway/go-sdk/compare/v1.0.3...v1.1.0) (2025-01-21)

### ✨ Features

* Update message roles to use typed constants and improve README examples ([a9107d2](https://github.com/inference-gateway/go-sdk/commit/a9107d27004a46f6b681aadd0a491b40aa5bf005))

## [1.0.3](https://github.com/inference-gateway/go-sdk/compare/v1.0.2...v1.0.3) (2025-01-21)

### ♻️ Improvements

* Update GenerateContent method to accept messages slice instead of prompt string ([b9e6e03](https://github.com/inference-gateway/go-sdk/commit/b9e6e03f3cc300c648f343e43d7ac847b65c63a9))

## [1.0.2](https://github.com/inference-gateway/go-sdk/compare/v1.0.1...v1.0.2) (2025-01-21)

### 🐛 Bug Fixes

* Update go.mod ([338036a](https://github.com/inference-gateway/go-sdk/commit/338036a8dd3fd8136a8d79c77169561e4defb5e7))

### 🔧 Miscellaneous

* Add .editorconfig for consistent coding styles across files ([0a4083f](https://github.com/inference-gateway/go-sdk/commit/0a4083f4800b1a0ab3a367fe3dbb6ddd16c828b3))

## [1.0.1](https://github.com/inference-gateway/go-sdk/compare/v1.0.0...v1.0.1) (2025-01-21)

### 📚 Documentation

* Add CONTRIBUTING.md with guidelines for contributing to the SDK ([e890b5d](https://github.com/inference-gateway/go-sdk/commit/e890b5ddc9b81a8a776bc1067e32de4677f0f4c6))

## 1.0.0 (2025-01-21)

### ♻️ Improvements

* Enhance SDK client structure and error handling for model listing and content generation ([095a453](https://github.com/inference-gateway/go-sdk/commit/095a4532bae14e65a3e4779be628f0f8727dbc13))

### 👷 CI

* Add a basic go ci ([124fb0f](https://github.com/inference-gateway/go-sdk/commit/124fb0f7636f99c0688c46fae0484f7db68163d0))
* Add GitHub Release workflow with semantic release integration ([127a22e](https://github.com/inference-gateway/go-sdk/commit/127a22ef65f61a2a84d943eab9aac25787f358cb))
* Rename workflow and job for clarity ([f23b09a](https://github.com/inference-gateway/go-sdk/commit/f23b09a9f27436b99cf6bebf260a9be2ece179f8))
* Update CI workflow to use Ubuntu 24.04 ([b2554b7](https://github.com/inference-gateway/go-sdk/commit/b2554b783348a75244c1f0392870889e4c47f5a9))

### 📚 Documentation

* Update README to reflect correct GitHub repository links and add missing health check, supported providers, and contributing sections ([5fb830b](https://github.com/inference-gateway/go-sdk/commit/5fb830b5ed03ad739b97dcc11f5ccf95a166cc2b))
* Update README with installation and usage instructions for the Inference Gateway Go SDK ([f0c0d1e](https://github.com/inference-gateway/go-sdk/commit/f0c0d1e70e3de1ad09734a0e9337bbbfd76c9bf0))

### 🔧 Miscellaneous

* Add initial devcontainer setup with Go SDK and configuration files ([d0ced5c](https://github.com/inference-gateway/go-sdk/commit/d0ced5c3eb03830f85ec0c6d42cfbd08186dd6ca))
* Create LICENSE ([87fa973](https://github.com/inference-gateway/go-sdk/commit/87fa97331e07c63613c62de3a66660937c4e2dcb))
* Run task oas-download to get the latest OpenAPI spec ([ce46514](https://github.com/inference-gateway/go-sdk/commit/ce465144e9a4205d850beadba6e7012db6cfa923))
* Update devcontainer setup with Task and semantic-release integration ([b138601](https://github.com/inference-gateway/go-sdk/commit/b138601ce423c9d728e4c9636fff5f935c70e03c))
* Update README.md ([7fbcca2](https://github.com/inference-gateway/go-sdk/commit/7fbcca26261309ffc4b44789f5cfc958ebb403bc))
* Update README.md ([7d08ca5](https://github.com/inference-gateway/go-sdk/commit/7d08ca52cdfff2032e384410ea3502b2d97f0488))

### ✅ Miscellaneous

* Add unit tests for model listing, content generation, and health check ([bc4b8a9](https://github.com/inference-gateway/go-sdk/commit/bc4b8a90f11cbce002f3fe23adc0afa821fea315))
