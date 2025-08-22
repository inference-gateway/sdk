# Changelog

All notable changes to this project will be documented in this file.

## [1.12.0-rc.1](https://github.com/inference-gateway/sdk/compare/v1.11.1...v1.12.0-rc.1) (2025-08-22)

### ✨ Features

* **retry:** Add configurable status codes and callback mechanism ([1d3cefa](https://github.com/inference-gateway/sdk/commit/1d3cefa4a8264fea954267cc8abc29c5a90e2f17))
* **retry:** Add retry logic with exponential backoff for HTTP requests ([0d1a57a](https://github.com/inference-gateway/sdk/commit/0d1a57af0b55bcffa4b70f6e790f71707f521960))

### ♻️ Improvements

* **retry:** Clean up comments ([b5a32db](https://github.com/inference-gateway/sdk/commit/b5a32db07aa88e943ed893e0607b8aaba2cf29dc))
* **retry:** Remove comments from retryable status code tests for clarity ([5948350](https://github.com/inference-gateway/sdk/commit/59483508e16dc5e4dcb73ad6ce6aa5a01b73432f))
* **retry:** Remove redundant comments in isRetryableStatusCode function ([930dc15](https://github.com/inference-gateway/sdk/commit/930dc1577df93ab6e2238d969e42157a80c78c0f))

### 🐛 Bug Fixes

* **headers:** Remove redundant comment in TestWithHeaders ([468f356](https://github.com/inference-gateway/sdk/commit/468f35616933900d946c9e2dfa76d47c314aa792))

### 🔧 Miscellaneous

* **release:** 🔖 1.12.0-rc.1 [skip ci] ([d36f1ab](https://github.com/inference-gateway/sdk/commit/d36f1ab08f77789e67db018f825b9ae5965793c8))
* Remove redundant comments ([fbbe49f](https://github.com/inference-gateway/sdk/commit/fbbe49fc51cd74db5f9414655bef0558c5e1b0a4))
* Testing the release ([#22](https://github.com/inference-gateway/sdk/issues/22)) ([05b9687](https://github.com/inference-gateway/sdk/commit/05b9687049680fafeeb95d153b2ebc9b7e2be055))

## [1.12.0-rc.1](https://github.com/inference-gateway/sdk/compare/v1.11.1...v1.12.0-rc.1) (2025-08-22)

### ✨ Features

* **retry:** Add configurable status codes and callback mechanism ([1d3cefa](https://github.com/inference-gateway/sdk/commit/1d3cefa4a8264fea954267cc8abc29c5a90e2f17))
* **retry:** Add retry logic with exponential backoff for HTTP requests ([0d1a57a](https://github.com/inference-gateway/sdk/commit/0d1a57af0b55bcffa4b70f6e790f71707f521960))

### ♻️ Improvements

* **retry:** Clean up comments ([b5a32db](https://github.com/inference-gateway/sdk/commit/b5a32db07aa88e943ed893e0607b8aaba2cf29dc))

### 🐛 Bug Fixes

* **headers:** Remove redundant comment in TestWithHeaders ([a6a6cbb](https://github.com/inference-gateway/sdk/commit/a6a6cbbc8f69eb6f5cfe8b1d2fcc1e33527dc9e9))

### 🔧 Miscellaneous

* Remove redundant comments ([fbbe49f](https://github.com/inference-gateway/sdk/commit/fbbe49fc51cd74db5f9414655bef0558c5e1b0a4))
* Testing the release ([#22](https://github.com/inference-gateway/sdk/issues/22)) ([05b9687](https://github.com/inference-gateway/sdk/commit/05b9687049680fafeeb95d153b2ebc9b7e2be055))

## [1.12.0-rc.1](https://github.com/inference-gateway/sdk/compare/v1.11.1...v1.12.0-rc.1) (2025-08-22)

### ✨ Features

* **retry:** Add configurable status codes and callback mechanism ([1d3cefa](https://github.com/inference-gateway/sdk/commit/1d3cefa4a8264fea954267cc8abc29c5a90e2f17))
* **retry:** Add retry logic with exponential backoff for HTTP requests ([0d1a57a](https://github.com/inference-gateway/sdk/commit/0d1a57af0b55bcffa4b70f6e790f71707f521960))

### ♻️ Improvements

* **retry:** Clean up comments ([b5a32db](https://github.com/inference-gateway/sdk/commit/b5a32db07aa88e943ed893e0607b8aaba2cf29dc))
* **tests:** Remove unnecessary blank line in TestIsRetryableStatusCode ([223ab26](https://github.com/inference-gateway/sdk/commit/223ab2679d3cfdc3080f3f1e4d5588cccf573bf9))

### 🔧 Miscellaneous

* Remove redundant comments ([fbbe49f](https://github.com/inference-gateway/sdk/commit/fbbe49fc51cd74db5f9414655bef0558c5e1b0a4))

## [1.11.1](https://github.com/inference-gateway/sdk/compare/v1.11.0...v1.11.1) (2025-08-20)

### 📚 Documentation

* **examples:** Add ReasoningContent field support in streaming examples ([#18](https://github.com/inference-gateway/sdk/issues/18)) ([0212d7d](https://github.com/inference-gateway/sdk/commit/0212d7df0c3e4ca5d94e9875f32676934a1f43eb)), closes [#17](https://github.com/inference-gateway/sdk/issues/17)

## [1.11.0](https://github.com/inference-gateway/sdk/compare/v1.10.0...v1.11.0) (2025-08-09)

### ✨ Features

* **providers:** Add Mistral AI provider support ([#16](https://github.com/inference-gateway/sdk/issues/16)) ([202cd9a](https://github.com/inference-gateway/sdk/commit/202cd9aa06e4f9d3c6cae3ab35803628df76e72a)), closes [#15](https://github.com/inference-gateway/sdk/issues/15)

### 👷 CI

* Add Claude Code GitHub Workflow ([#13](https://github.com/inference-gateway/sdk/issues/13)) ([8bb2237](https://github.com/inference-gateway/sdk/commit/8bb2237836de996966c9bb4f6f380b82fefd918c))

### 📚 Documentation

* Simplify CLAUDE.md by consolidating commands and development guidelines ([50b02c5](https://github.com/inference-gateway/sdk/commit/50b02c5c6d2d9f504643b6d926bab20f0b3e8aac))

### 🔧 Miscellaneous

* Add issue templates for bug reports, feature requests, and refactor requests ([#14](https://github.com/inference-gateway/sdk/issues/14)) ([9d3371b](https://github.com/inference-gateway/sdk/commit/9d3371bc4d275aa6228748275450623ce2bd68db))

### 📦 Miscellaneous

* Add initial Flox configuration files for SDK environment ([fafcba4](https://github.com/inference-gateway/sdk/commit/fafcba4460da88bd20dd98566f51dc7b3988441f))

## [1.10.0](https://github.com/inference-gateway/sdk/compare/v1.9.1...v1.10.0) (2025-07-26)

### ✨ Features

* Add Google Provider ([#12](https://github.com/inference-gateway/sdk/issues/12)) ([f963677](https://github.com/inference-gateway/sdk/commit/f9636774d2e535ac0ebf22ffca1ea1c830b5f009))

### 📚 Documentation

* Add Google provider to supported LLM providers list in README ([b4f6adf](https://github.com/inference-gateway/sdk/commit/b4f6adf7d642d18b5553c1fa4c889f71075a59d1))

## [1.9.1](https://github.com/inference-gateway/sdk/compare/v1.9.0...v1.9.1) (2025-07-19)

### 📚 Documentation

* Add middleware options section to README with usage examples ([47a703e](https://github.com/inference-gateway/sdk/commit/47a703ee449451aa77b06228bd2905a0e47d7a21))

### ✅ Miscellaneous

* Enhance middleware options handling and add comprehensive tests ([4e09828](https://github.com/inference-gateway/sdk/commit/4e098286eb9f8cb9f6fb61c0ab1ddb8af40c0196))

## [1.9.0](https://github.com/inference-gateway/sdk/compare/v1.8.3...v1.9.0) (2025-07-19)

### ✨ Features

* Add middleware bypass example and support for middleware options ([4d78ae9](https://github.com/inference-gateway/sdk/commit/4d78ae9ce4eda57aafa6a1517fa9226e82a3db03))

## [1.8.3](https://github.com/inference-gateway/sdk/compare/v1.8.2...v1.8.3) (2025-06-20)

### 📚 Documentation

* Improve README ([ece2d30](https://github.com/inference-gateway/sdk/commit/ece2d304cc6c899a1e02a99c96371c6dd54b7f87))

## [1.8.2](https://github.com/inference-gateway/sdk/compare/v1.8.1...v1.8.2) (2025-06-20)

### 👷 CI

* Update Go version to 1.24 and add GolangCI Lint step in CI workflow ([4e89580](https://github.com/inference-gateway/sdk/commit/4e8958001d2d639167d82e9cb0333d6f5c634933))

### 📚 Documentation

* Add examples section to README and update MCP Tools link ([2260a2b](https://github.com/inference-gateway/sdk/commit/2260a2be8f60f0f082c9e6eea7b42ae01b72d7b7))
* Add real time streaming Tools Agent with weather and calculator functionalities ([4cd2ca4](https://github.com/inference-gateway/sdk/commit/4cd2ca445b5564b43e83d65911fd9305ac5fccf9))

### 🔧 Miscellaneous

* Update date in Copilot instructions to reflect current date for up to date research ([0260471](https://github.com/inference-gateway/sdk/commit/0260471eb7793224ce72ce45198f7864c769b71c))
* Upgrade Go version to 1.24 and update dependencies to v1.8.1 ([78b263b](https://github.com/inference-gateway/sdk/commit/78b263b1d887750efee58065e89190feb2714e55))

## [1.8.1](https://github.com/inference-gateway/sdk/compare/v1.8.0...v1.8.1) (2025-06-10)

### ♻️ Improvements

* Improve SDK - allow to pass http custom headers ([#10](https://github.com/inference-gateway/sdk/issues/10)) ([654cabc](https://github.com/inference-gateway/sdk/commit/654cabc4ed137b5cfca2dcc3a1d008e8e1be0042))

## [1.8.0](https://github.com/inference-gateway/sdk/compare/v1.7.1...v1.8.0) (2025-05-26)

### ✨ Features

* Implement List MCP Tools ([#9](https://github.com/inference-gateway/sdk/issues/9)) ([c8020e1](https://github.com/inference-gateway/sdk/commit/c8020e18400b635306efb64c41436b1870583c9d))

## [1.7.1](https://github.com/inference-gateway/sdk/compare/v1.7.0...v1.7.1) (2025-05-22)

### 🐛 Bug Fixes

* Improve error handling for API errors in content generation and streaming ([4e63a87](https://github.com/inference-gateway/sdk/commit/4e63a87a9ae9d4009c2c9a7866de6a974423c2d2))

### 🔧 Miscellaneous

* Add Copilot instructions and improve error handling in content stream generation ([205ed95](https://github.com/inference-gateway/sdk/commit/205ed9574752cca065817bdc32e2de4f81c2f39c))
* Add task completion support to zsh configuration ([ca452af](https://github.com/inference-gateway/sdk/commit/ca452af3b66769c1ee986c0c9d572a01c0fc05e0))
* Enhance Copilot instructions for pull request and test generation ([f8590cb](https://github.com/inference-gateway/sdk/commit/f8590cbaf04b65b77c89656caa689d7869ea01f7))
* update dependencies to latest versions in go.mod and go.sum ([36c39f3](https://github.com/inference-gateway/sdk/commit/36c39f364203f572b343fb26f97b4838dbae444d))

### 📦 Miscellaneous

* Update devcontainer configuration and remove unused files ([127bde4](https://github.com/inference-gateway/sdk/commit/127bde435b61248a36e1db590e1cc3d14baa1b2b))

## [1.7.0](https://github.com/inference-gateway/sdk/compare/v1.6.0...v1.7.0) (2025-04-30)

### ✨ Features

* Update API spec and add reasoning format support ([#6](https://github.com/inference-gateway/sdk/issues/6)) ([f889151](https://github.com/inference-gateway/sdk/commit/f889151cac10656c12fae0240bca1f91ae067e1a)), closes [#7](https://github.com/inference-gateway/sdk/issues/7)

### 📚 Documentation

* Add CLAUDE.md with development guidelines ([c451ec4](https://github.com/inference-gateway/sdk/commit/c451ec4e48b10f7fb24899bc57300909943c889e))

### 🔧 Miscellaneous

* Add Claude Code CLI to development container ([2e716eb](https://github.com/inference-gateway/sdk/commit/2e716eb08edc387b6a720626c825fd32cc829fdc))

## [1.7.0-rc.2](https://github.com/inference-gateway/sdk/compare/v1.7.0-rc.1...v1.7.0-rc.2) (2025-04-30)

### ✨ Features

-   Add Reasoning and ReasoningContent fields to ChatCompletionStreamResponseDelta ([1f0143e](https://github.com/inference-gateway/sdk/commit/1f0143e9e8837042c6f0b6c01e89876a9f05475a))

## [1.7.0-rc.1](https://github.com/inference-gateway/sdk/compare/v1.6.0...v1.7.0-rc.1) (2025-04-30)

### ✨ Features

-   Add WithOptions method and support for reasoning formats ([d929814](https://github.com/inference-gateway/sdk/commit/d9298147f955e77de9ae248ecd7d058c91b0e30c))
-   Update to latest OpenAPI spec and model ID format ([537cd1f](https://github.com/inference-gateway/sdk/commit/537cd1ffe3fe70c26ec6e740f48880ad9b9c750c))

### 📚 Documentation

-   Add CLAUDE.md with development guidelines ([c451ec4](https://github.com/inference-gateway/sdk/commit/c451ec4e48b10f7fb24899bc57300909943c889e))
-   Add information about ReasoningFormat and Reasoning fields ([f8d667d](https://github.com/inference-gateway/sdk/commit/f8d667df8bee90bbb987f66d711b2bb35932cb32))
-   Update CLAUDE.md with latest API conventions ([7845cf8](https://github.com/inference-gateway/sdk/commit/7845cf8afb539e75c3857a892cc962811f56b042))
-   Update README.md to reflect API changes ([7b6e315](https://github.com/inference-gateway/sdk/commit/7b6e315df08e64ee938e56abca60176f54d2ccfb))

### 🔧 Miscellaneous

-   Add binary files to .gitignore ([97cdbca](https://github.com/inference-gateway/sdk/commit/97cdbcaadf8cbf73f37acd1a40d636a572318c31))
-   Add Claude Code CLI to development container ([2e716eb](https://github.com/inference-gateway/sdk/commit/2e716eb08edc387b6a720626c825fd32cc829fdc))
-   Ensure newline at end of .gitignore for consistency ([77b1511](https://github.com/inference-gateway/sdk/commit/77b1511eedbdf0423646434d5cc10eab1af621fa))
-   Update .gitignore to properly ignore binaries ([35a74e9](https://github.com/inference-gateway/sdk/commit/35a74e93e242ff049a97e97caf50a0303c829e4c))

## [1.6.0](https://github.com/inference-gateway/sdk/compare/v1.5.1...v1.6.0) (2025-03-25)

### ✨ Features

-   Add Deepseek provider ([#5](https://github.com/inference-gateway/sdk/issues/5)) ([28ecd4a](https://github.com/inference-gateway/sdk/commit/28ecd4a4fb230a610c05ff491fbbfc5b829309c4))

### 📚 Documentation

-   Update comment for Tools field in ClientOptions for clarity ([299ab0f](https://github.com/inference-gateway/sdk/commit/299ab0fa6737cb32190cb9c4f210badd4f435513))

## [1.5.1](https://github.com/inference-gateway/sdk/compare/v1.5.0...v1.5.1) (2025-03-25)

### ♻️ Improvements

-   Update NewClient to accept ClientOptions for improved configuration ([#4](https://github.com/inference-gateway/sdk/issues/4)) ([efd181f](https://github.com/inference-gateway/sdk/commit/efd181fcbb320356946aba8a5982353d830f1c85))

### 📚 Documentation

-   Add Tool-Use section and reorganize Health Check in README ([2e5d9e6](https://github.com/inference-gateway/sdk/commit/2e5d9e6004fce4c824cea0278318e664349708bf))
-   Update example README to correct link and format output section ([2d8abff](https://github.com/inference-gateway/sdk/commit/2d8abfffab13c8f84a3e81f3339dcc98617c92e0))

## [1.5.0](https://github.com/inference-gateway/sdk/compare/v1.4.1...v1.5.0) (2025-03-24)

### ✨ Features

-   Make this SDK OpenAI compatible ([#2](https://github.com/inference-gateway/sdk/issues/2)) ([3181988](https://github.com/inference-gateway/sdk/commit/318198864f14592ac5910c459417535569202737))

## [1.5.0-rc.2](https://github.com/inference-gateway/sdk/compare/v1.5.0-rc.1...v1.5.0-rc.2) (2025-03-24)

### ✨ Features

-   Add example applications for content generation, model listing, streaming, and tools calling ([c070421](https://github.com/inference-gateway/sdk/commit/c0704213dd2652baaf8d345571a8d06144253763))

### 🐛 Bug Fixes

-   **sdk:** Update API endpoint to remove versioning from URL ([19cd817](https://github.com/inference-gateway/sdk/commit/19cd81731c839ebcdcf33f8a68ce7330f7eae074))

### 📚 Documentation

-   **examples:** Add content generation example using Inference Gateway SDK ([bb62e29](https://github.com/inference-gateway/sdk/commit/bb62e29cf847028fe7a7e00f0de0251203fa88a8))
-   **examples:** Add example applications and configuration files for Inference Gateway SDK usage ([1d1b244](https://github.com/inference-gateway/sdk/commit/1d1b244c5f39f6ab884fa995234d392cdcdba9a1))
-   **examples:** Ensure model listing example works ([a27cba4](https://github.com/inference-gateway/sdk/commit/a27cba49b99c919c88c8f2f2af7db28763f848d1))
-   **examples:** Remove redundant nonsense statement generated by the LLM ([b04ef3f](https://github.com/inference-gateway/sdk/commit/b04ef3f9d9a69e951009c81fc070a82a8ef14e1a))
-   **examples:** Temporarily comment it out ([fb53e89](https://github.com/inference-gateway/sdk/commit/fb53e89fcdb109e3df176a1ceb05e80ed787095d))
-   **examples:** Update example section titles for clarity and consistency ([522a8b1](https://github.com/inference-gateway/sdk/commit/522a8b1da5cf8a503e62472c00b3d5f731fdd120))

## [1.5.0-rc.1](https://github.com/inference-gateway/sdk/compare/v1.4.1...v1.5.0-rc.1) (2025-03-23)

### ✨ Features

-   Add task to generate code from OpenAPI specification and define SSEvent schema ([17dfa8f](https://github.com/inference-gateway/sdk/commit/17dfa8f56ca9180d2a6233973ac3de56f8456830))

### ♻️ Improvements

-   Add error handling for JSON encoding in test cases ([ad38d99](https://github.com/inference-gateway/sdk/commit/ad38d993de8da1f020a21c288fea14feafef4e7c))
-   Clarify TODO comment regarding error event type in GenerateContentStream method ([dff7c6b](https://github.com/inference-gateway/sdk/commit/dff7c6bb559bd97aaefa09d0d851a35e62b0e8c7))
-   Download OpenAPI spec from inference-gateway ([16d8ddc](https://github.com/inference-gateway/sdk/commit/16d8ddc310416815513f8384ba080f6063f21690))
-   Enhance GenerateContent and GenerateContentStream methods for improved error handling and streaming support ([19d5e98](https://github.com/inference-gateway/sdk/commit/19d5e98b5866844a047971b2a909546f8ffdb7ad))
-   Implement ListModels and ListProviderModels methods ([61a1e06](https://github.com/inference-gateway/sdk/commit/61a1e06c5365e00d18ad8b4c36dece87d88d12ae))
-   Make it possible to have a RC in the release workflow ([26c8f0b](https://github.com/inference-gateway/sdk/commit/26c8f0bf08a761ae70378240a444e1b25c9d93e9))
-   Place the comment for clarity on the right attribute ([1a75bcb](https://github.com/inference-gateway/sdk/commit/1a75bcba82f3adeb2a8681dc57a854bfd97f4ed3))
-   Rename Providers to Provider and update related schema references ([c114701](https://github.com/inference-gateway/sdk/commit/c114701551032cecdd6a214ba33a54ef3e265e78))
-   Run task generate ([6655266](https://github.com/inference-gateway/sdk/commit/66552664fc8968ccbd1ae8b44e713e5998f600cd))
-   Update GenerateContent method to return CreateChatCompletionResponse ([63d5585](https://github.com/inference-gateway/sdk/commit/63d5585dce897a7c54df48d4686bfd340eb73a09))
-   Update ListModels and ListProviderModels methods to return pointers to ListModelsResponse ([2534da0](https://github.com/inference-gateway/sdk/commit/2534da0e1ab9fe29b8ffd25c75e9a64cd8378915))
-   Update OpenAPI schema and SDK to support CreateChatCompletionResponse structure ([5f60537](https://github.com/inference-gateway/sdk/commit/5f60537527c921729c395afcd45244470541c156))
-   Update output filename for generated Go types to generated_types.go ([a8bc1b6](https://github.com/inference-gateway/sdk/commit/a8bc1b6517a6b334c6b301d0375ac318907c8902))
-   Update README for improved client usage and error handling ([911e8e5](https://github.com/inference-gateway/sdk/commit/911e8e5f76e1906c816ae2f80e722a3a683aed1c))
-   Update task description to clarify code generation for Go types ([886b130](https://github.com/inference-gateway/sdk/commit/886b1306172df04653404a2711cacaac3518996b))

### 📚 Documentation

-   Download the latest OpenAPI specification ([ee60b31](https://github.com/inference-gateway/sdk/commit/ee60b311026d28d16f69dc7d5faee5cc505896b0))

## [1.4.1](https://github.com/inference-gateway/sdk/compare/v1.4.0...v1.4.1) (2025-02-02)

### 📚 Documentation

-   Update README with the correct examples ([6b8268d](https://github.com/inference-gateway/sdk/commit/6b8268de350580fec09d2c36a92ad131929c2518))

## [1.4.0](https://github.com/inference-gateway/sdk/compare/v1.3.0...v1.4.0) (2025-02-02)

### ✨ Features

-   Add streaming content generation support to SDK ([6f4b4a0](https://github.com/inference-gateway/sdk/commit/6f4b4a046dbd011d3c7f04f0d28b6f5e96bcd73c))

### ♻️ Improvements

-   Add context parameter to SDK client methods for improved context management ([cac2f4e](https://github.com/inference-gateway/sdk/commit/cac2f4edd8494b84bd7007a7a1dd67646baabdf0))
-   Rename role constants for clarity in SDK message handling ([2782890](https://github.com/inference-gateway/sdk/commit/278289059f998d372c387dc85cb237b981253855))

### 📚 Documentation

-   Correct typo in Anthropic provider listing in README ([0195de7](https://github.com/inference-gateway/sdk/commit/0195de724f119eacaac8412dcf3b0a5fc92822bb))
-   Improve docs by having a docblocks with examples ([19f105d](https://github.com/inference-gateway/sdk/commit/19f105d528de3aa0950d48d473c222d66c0f9701))
-   Remove Google LLM provider from SDK and documentation ([8ba8ec0](https://github.com/inference-gateway/sdk/commit/8ba8ec0c1b5e6c8183f74a05a3ffffb3dbd5eaf0))

### 🔧 Miscellaneous

-   Cleanup ([f6547a8](https://github.com/inference-gateway/sdk/commit/f6547a8d4abb39540dcd4d3f34ef9556ad3a3d47))
-   **openapi:** Update documentation for SDK usage ([f84968d](https://github.com/inference-gateway/sdk/commit/f84968d694cc3b8b805670bef020b3a417520fdc))

## [1.3.0](https://github.com/inference-gateway/sdk/compare/v1.2.3...v1.3.0) (2025-01-28)

### ✨ Features

-   Add ListProviderModels to list a specific providers models ([1cf86fb](https://github.com/inference-gateway/sdk/commit/1cf86fb0debf58b54d3f5d76575fce3672c43f9e))
-   Add support for Anthropic LLM provider in SDK ([fa5772a](https://github.com/inference-gateway/sdk/commit/fa5772a23e9e8d358081b48479460a24aceb2653))

### 📦 Improvements

-   Add lint task to run golangci-lint in Taskfile ([60403f1](https://github.com/inference-gateway/sdk/commit/60403f1a746fb481fc04d7980025aeef91666f2c))

### 📚 Documentation

-   **openapi:** Update OpenAPI spec - download the latest one from inference-gateway ([2056147](https://github.com/inference-gateway/sdk/commit/205614711a886998e7179fe4b50d16b2eacd65d7))

## [1.2.3](https://github.com/inference-gateway/sdk/compare/v1.2.2...v1.2.3) (2025-01-21)

### ♻️ Improvements

-   Change Client type to interface and implement clientImpl for SDK client methods ([003b660](https://github.com/inference-gateway/sdk/commit/003b66053fdd069ddb5ee1435c211aa9b53362c2))

### 🔧 Miscellaneous

-   Change semantic-release configurations build prefix to be listed under Improvements ([0ab0d02](https://github.com/inference-gateway/sdk/commit/0ab0d02c2011dfef1292ab4541ac8e5df206dd93))

## [1.2.2](https://github.com/inference-gateway/sdk/compare/v1.2.1...v1.2.2) (2025-01-21)

### 📦 Miscellaneous

-   Add VSCode task extension to development container configuration for easy execution of tasks ([3592065](https://github.com/inference-gateway/sdk/commit/35920650d54122d29716d489efbc35025c13eebf))

## [1.2.1](https://github.com/inference-gateway/sdk/compare/v1.2.0...v1.2.1) (2025-01-21)

### 🐛 Bug Fixes

-   Update repository references in configuration files and documentation ([39d0041](https://github.com/inference-gateway/sdk/commit/39d0041a7a2cdbac77083a40f086bf0c8a9cc5d8))

## [1.2.0](https://github.com/inference-gateway/go-sdk/compare/v1.1.0...v1.2.0) (2025-01-21)

### ✨ Features

-   Add Go tools installation and enhance Taskfile with tidy, test, and docs tasks ([9c27a1d](https://github.com/inference-gateway/go-sdk/commit/9c27a1d8db4aff95da3b03a56b4ed0e911ffe1bd))

### ♻️ Improvements

-   Update module name and README for improved clarity and documentation ([59ef21a](https://github.com/inference-gateway/go-sdk/commit/59ef21ae2d652707dd5b3d5a526cc8f2bcbdd355))

### 🐛 Bug Fixes

-   Correct spelling of "helpful" in README and update tests to reflect changes ([ba7a423](https://github.com/inference-gateway/go-sdk/commit/ba7a4237a9f523d5e6a0421cab62eaf56e59c258))

### 📚 Documentation

-   Enhance SDK documentation with detailed comments and examples for clarity ([1727352](https://github.com/inference-gateway/go-sdk/commit/17273529263efce5bac34473b8bec2fdde0d3760))

## [1.1.0](https://github.com/inference-gateway/go-sdk/compare/v1.0.3...v1.1.0) (2025-01-21)

### ✨ Features

-   Update message roles to use typed constants and improve README examples ([a9107d2](https://github.com/inference-gateway/go-sdk/commit/a9107d27004a46f6b681aadd0a491b40aa5bf005))

## [1.0.3](https://github.com/inference-gateway/go-sdk/compare/v1.0.2...v1.0.3) (2025-01-21)

### ♻️ Improvements

-   Update GenerateContent method to accept messages slice instead of prompt string ([b9e6e03](https://github.com/inference-gateway/go-sdk/commit/b9e6e03f3cc300c648f343e43d7ac847b65c63a9))

## [1.0.2](https://github.com/inference-gateway/go-sdk/compare/v1.0.1...v1.0.2) (2025-01-21)

### 🐛 Bug Fixes

-   Update go.mod ([338036a](https://github.com/inference-gateway/go-sdk/commit/338036a8dd3fd8136a8d79c77169561e4defb5e7))

### 🔧 Miscellaneous

-   Add .editorconfig for consistent coding styles across files ([0a4083f](https://github.com/inference-gateway/go-sdk/commit/0a4083f4800b1a0ab3a367fe3dbb6ddd16c828b3))

## [1.0.1](https://github.com/inference-gateway/go-sdk/compare/v1.0.0...v1.0.1) (2025-01-21)

### 📚 Documentation

-   Add CONTRIBUTING.md with guidelines for contributing to the SDK ([e890b5d](https://github.com/inference-gateway/go-sdk/commit/e890b5ddc9b81a8a776bc1067e32de4677f0f4c6))

## 1.0.0 (2025-01-21)

### ♻️ Improvements

-   Enhance SDK client structure and error handling for model listing and content generation ([095a453](https://github.com/inference-gateway/go-sdk/commit/095a4532bae14e65a3e4779be628f0f8727dbc13))

### 👷 CI

-   Add a basic go ci ([124fb0f](https://github.com/inference-gateway/go-sdk/commit/124fb0f7636f99c0688c46fae0484f7db68163d0))
-   Add GitHub Release workflow with semantic release integration ([127a22e](https://github.com/inference-gateway/go-sdk/commit/127a22ef65f61a2a84d943eab9aac25787f358cb))
-   Rename workflow and job for clarity ([f23b09a](https://github.com/inference-gateway/go-sdk/commit/f23b09a9f27436b99cf6bebf260a9be2ece179f8))
-   Update CI workflow to use Ubuntu 24.04 ([b2554b7](https://github.com/inference-gateway/go-sdk/commit/b2554b783348a75244c1f0392870889e4c47f5a9))

### 📚 Documentation

-   Update README to reflect correct GitHub repository links and add missing health check, supported providers, and contributing sections ([5fb830b](https://github.com/inference-gateway/go-sdk/commit/5fb830b5ed03ad739b97dcc11f5ccf95a166cc2b))
-   Update README with installation and usage instructions for the Inference Gateway Go SDK ([f0c0d1e](https://github.com/inference-gateway/go-sdk/commit/f0c0d1e70e3de1ad09734a0e9337bbbfd76c9bf0))

### 🔧 Miscellaneous

-   Add initial devcontainer setup with Go SDK and configuration files ([d0ced5c](https://github.com/inference-gateway/go-sdk/commit/d0ced5c3eb03830f85ec0c6d42cfbd08186dd6ca))
-   Create LICENSE ([87fa973](https://github.com/inference-gateway/go-sdk/commit/87fa97331e07c63613c62de3a66660937c4e2dcb))
-   Run task oas-download to get the latest OpenAPI spec ([ce46514](https://github.com/inference-gateway/go-sdk/commit/ce465144e9a4205d850beadba6e7012db6cfa923))
-   Update devcontainer setup with Task and semantic-release integration ([b138601](https://github.com/inference-gateway/go-sdk/commit/b138601ce423c9d728e4c9636fff5f935c70e03c))
-   Update README.md ([7fbcca2](https://github.com/inference-gateway/go-sdk/commit/7fbcca26261309ffc4b44789f5cfc958ebb403bc))
-   Update README.md ([7d08ca5](https://github.com/inference-gateway/go-sdk/commit/7d08ca52cdfff2032e384410ea3502b2d97f0488))

### ✅ Miscellaneous

-   Add unit tests for model listing, content generation, and health check ([bc4b8a9](https://github.com/inference-gateway/go-sdk/commit/bc4b8a90f11cbce002f3fe23adc0afa821fea315))
