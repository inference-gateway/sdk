# Changelog

All notable changes to this project will be documented in this file.

## [1.20.1](https://github.com/inference-gateway/sdk/compare/v1.20.0...v1.20.1) (2026-07-10)

### 🐛 Bug Fixes

* preserve HTTP error when context cancelled during backoff ([#119](https://github.com/inference-gateway/sdk/issues/119)) ([6998109](https://github.com/inference-gateway/sdk/commit/699810956c0d841bfb99d6ce75d70137899fe6ff))
* prevent stream reader goroutine leak on ctx cancel ([#121](https://github.com/inference-gateway/sdk/issues/121)) ([8c8ff2d](https://github.com/inference-gateway/sdk/commit/8c8ff2def8ae67c65749e0250b0f03ad683bc402))
* prevent stream reader panic on body close error ([#120](https://github.com/inference-gateway/sdk/issues/120)) ([f107881](https://github.com/inference-gateway/sdk/commit/f107881515b988437418a86341f61474a0b571ad)), closes [#116](https://github.com/inference-gateway/sdk/issues/116)

### 👷 CI

* **claude:** centralize claude.yml via reusable workflow ([#114](https://github.com/inference-gateway/sdk/issues/114)) ([3ce2ac8](https://github.com/inference-gateway/sdk/commit/3ce2ac80bf1c74c7c1fdd3df7faf0897a2bd68a3))
* **infer:** centralize infer.yml via reusable workflow ([#111](https://github.com/inference-gateway/sdk/issues/111)) ([61a6a2c](https://github.com/inference-gateway/sdk/commit/61a6a2c3d36d93bd266c1e2467b78cb641bb69fa))
* **infer:** centralize infer.yml via reusable workflow ([#112](https://github.com/inference-gateway/sdk/issues/112)) ([e3d3acc](https://github.com/inference-gateway/sdk/commit/e3d3acc2dc36f7d1ef919358c63886e9f4be2816))
* restrict default workflow token permissions to contents: read ([#110](https://github.com/inference-gateway/sdk/issues/110)) ([e67a0f2](https://github.com/inference-gateway/sdk/commit/e67a0f2d7b4afd5305fea482cac13818e410ca5b))

### 🔧 Miscellaneous

* **deps:** bump claude-code-action v1.0.165 -> v1.0.169 ([#113](https://github.com/inference-gateway/sdk/issues/113)) ([b49dec6](https://github.com/inference-gateway/sdk/commit/b49dec676744c24de1f1375550111d038007dfce))
* **deps:** bump infer CLI v0.137.0 -> v0.138.0, infer-action v0.29.0 -> v0.30.1 ([#109](https://github.com/inference-gateway/sdk/issues/109)) ([4acac0a](https://github.com/inference-gateway/sdk/commit/4acac0adc15b3eec275d01e9a59b20ad9c2c14bb))
* **deps:** bump infer CLI v0.138.0 -> v0.141.0 ([#115](https://github.com/inference-gateway/sdk/issues/115)) ([fea1e80](https://github.com/inference-gateway/sdk/commit/fea1e808cec2887c42119296066777b4516616c3))

## [1.20.0](https://github.com/inference-gateway/sdk/compare/v1.19.0...v1.20.0) (2026-07-08)

### ✨ Features

* add Responses API support with zai provider ([#107](https://github.com/inference-gateway/sdk/issues/107)) ([71128ab](https://github.com/inference-gateway/sdk/commit/71128ab3f953e7a872c1c1e15abb1a7369fa518b))

### 👷 CI

* **deps:** bump golang.org/x/net in /examples/middleware-bypass ([#104](https://github.com/inference-gateway/sdk/issues/104)) ([9fae726](https://github.com/inference-gateway/sdk/commit/9fae726650a5bb3e748619cb0e4c5cb4604c9ff4))
* **deps:** bump inference-gateway/.github/.github/workflows/claude.yml ([#106](https://github.com/inference-gateway/sdk/issues/106)) ([c3b667b](https://github.com/inference-gateway/sdk/commit/c3b667b3479deb09f5bab050f6284e035f96f1d3))
* **release:** update semantic release and plugins to latest versions with local installation ([2b64b79](https://github.com/inference-gateway/sdk/commit/2b64b79b8ca7a842de66959b484567ea85006827))

### 🔧 Miscellaneous

* **deps:** bump claude-code 2.1.177 -> 2.1.197, claude-code-action v1.0.161 -> v1.0.165 ([#100](https://github.com/inference-gateway/sdk/issues/100)) ([dc226c9](https://github.com/inference-gateway/sdk/commit/dc226c9d56c4f356dcf088002a6de010a72bf515))
* **deps:** bump claude-code 2.1.197 -> 2.1.201 ([#101](https://github.com/inference-gateway/sdk/issues/101)) ([a6960b6](https://github.com/inference-gateway/sdk/commit/a6960b667b10beb8060615c67922702a15ce31a3))
* **deps:** bump infer CLI v0.130.1 -> v0.133.0, infer-action v0.24.0 -> v0.26.0 ([#102](https://github.com/inference-gateway/sdk/issues/102)) ([74b8fc9](https://github.com/inference-gateway/sdk/commit/74b8fc9d3ae4fff16e61409cdecd25b819489b1e))
* **deps:** bump infer CLI v0.133.0 -> v0.133.1, infer-action v0.26.0 -> v0.27.1 ([#103](https://github.com/inference-gateway/sdk/issues/103)) ([9ea8113](https://github.com/inference-gateway/sdk/commit/9ea8113a541db78a09007c767341b944cde717a7))
* **deps:** bump infer CLI v0.133.1 -> v0.137.0, infer-action v0.27.1 -> v0.29.0 ([#105](https://github.com/inference-gateway/sdk/issues/105)) ([83b35c5](https://github.com/inference-gateway/sdk/commit/83b35c57d71628567efc12e03d230322c85ab237))
* **deps:** update dependencies and add go-task ([098f686](https://github.com/inference-gateway/sdk/commit/098f6862128f9570bafb99616afe9490cc303202))
* remove deprecated configuration and shortcut files ([238997b](https://github.com/inference-gateway/sdk/commit/238997b18809e18d0f58791f1553feae81cf6dd6))

## [1.19.0](https://github.com/inference-gateway/sdk/compare/v1.18.1...v1.19.0) (2026-07-05)

## [1.18.1](https://github.com/inference-gateway/sdk/compare/v1.18.0...v1.18.1) (2026-06-18)

### 👷 CI

* **deps:** upgrade actions/checkout from v6.0.3 to v7.0.0 across workflows ([bea800e](https://github.com/inference-gateway/sdk/commit/bea800eb148b6340117ed9166dfed6be07fcc228))

### 🔧 Miscellaneous

* **deps:** update schema version and codex version to 0.139.0 in manifest files ([80aa77e](https://github.com/inference-gateway/sdk/commit/80aa77e25e384df24020ae6223b8fe190c09b813))
* update CI workflow to include permissions ([ff64b16](https://github.com/inference-gateway/sdk/commit/ff64b16410ba0a520135fb374d12816fb74cc958))

## [1.18.0](https://github.com/inference-gateway/sdk/compare/v1.17.0...v1.18.0) (2026-06-18)

### ✨ Features

* regenerate SDK types from updated openapi.yaml ([#85](https://github.com/inference-gateway/sdk/issues/85)) ([a030c19](https://github.com/inference-gateway/sdk/commit/a030c191fbfeabfbc4046b4187eea026a68dccbf)), closes [schemas#71](https://github.com/inference-gateway/schemas/issues/71)

### 👷 CI

* **deps:** bump the github-actions group with 2 updates ([#81](https://github.com/inference-gateway/sdk/issues/81)) ([d66cdeb](https://github.com/inference-gateway/sdk/commit/d66cdeb58222623cc412249563c7721be8012b5d))
* **infer:** centralize infer.yml + sync .infer config ([#83](https://github.com/inference-gateway/sdk/issues/83)) ([8fff409](https://github.com/inference-gateway/sdk/commit/8fff409eb525c084737351df403522aef0ddfef6))

### 🔧 Miscellaneous

* **deps:** bump claude-code 2.1.161 -> 2.1.170, claude-code-action v1.0.135 -> v1.0.142 ([#79](https://github.com/inference-gateway/sdk/issues/79)) ([889e3d9](https://github.com/inference-gateway/sdk/commit/889e3d9cb70011aab89f4abe88661059ba2774ab))
* **deps:** bump claude-code 2.1.170 -> 2.1.177, claude-code-action v1.0.142 -> v1.0.150 ([#82](https://github.com/inference-gateway/sdk/issues/82)) ([9a026c7](https://github.com/inference-gateway/sdk/commit/9a026c749c27df585da96bc27724018605eef4b8))
* **deps:** bump infer CLI v0.121.0 -> v0.121.1, infer-action v0.12.1 -> v0.13.1 ([#80](https://github.com/inference-gateway/sdk/issues/80)) ([9a1fc16](https://github.com/inference-gateway/sdk/commit/9a1fc163b8a1a90fa922a75165956b0ae1e84f45))

## [1.17.0](https://github.com/inference-gateway/sdk/compare/v1.16.4...v1.17.0) (2026-06-11)

### ✨ Features

* add minimax to provider enum ([#78](https://github.com/inference-gateway/sdk/issues/78)) ([3f5bd68](https://github.com/inference-gateway/sdk/commit/3f5bd68fab5cd76e41b1c4dde829abb176aa82c8))

### 👷 CI

* centralize claude.yml via reusable workflow ([#57](https://github.com/inference-gateway/sdk/issues/57)) ([75a92af](https://github.com/inference-gateway/sdk/commit/75a92af9252344d64d0b80791fb7867d080dce20))
* centralize claude.yml via reusable workflow ([#58](https://github.com/inference-gateway/sdk/issues/58)) ([ba399f0](https://github.com/inference-gateway/sdk/commit/ba399f0a3f594515019cb14c90cb11b4a97cb17f))
* centralize claude.yml via reusable workflow ([#59](https://github.com/inference-gateway/sdk/issues/59)) ([da8d999](https://github.com/inference-gateway/sdk/commit/da8d9993f1081e22c572076870c866bb488d2930))
* centralize claude.yml via reusable workflow ([#75](https://github.com/inference-gateway/sdk/issues/75)) ([c7af56c](https://github.com/inference-gateway/sdk/commit/c7af56c604928c7c657dc7a7f5bb4589cce01a1b))
* centralize infer.yml + bump infer CLI and sync .infer config ([#62](https://github.com/inference-gateway/sdk/issues/62)) ([12de81b](https://github.com/inference-gateway/sdk/commit/12de81bc2b532f5d79d1091adc0f1b83e5ce7b3c))
* centralize infer.yml + sync .infer config ([#61](https://github.com/inference-gateway/sdk/issues/61)) ([9252034](https://github.com/inference-gateway/sdk/commit/92520349d521e8c975cdc58dbe04faba6fa2104f))
* centralize infer.yml via reusable workflow ([#60](https://github.com/inference-gateway/sdk/issues/60)) ([fd172d7](https://github.com/inference-gateway/sdk/commit/fd172d780f8a5f25779e3aaf5b0364f5ea1a4025))
* **claude:** Add maintainer skill ([6daeb97](https://github.com/inference-gateway/sdk/commit/6daeb975f560a8dece5f3cbadc6622cdf441b4f5))
* **claude:** change effort to max ([4b68110](https://github.com/inference-gateway/sdk/commit/4b6811028f41ecb88c77bbcfe16f391abe08e6f5))
* **claude:** download all maintainer skill assets ([48e51b4](https://github.com/inference-gateway/sdk/commit/48e51b4d4ba911e14000df4a246dfdbe15d1d290))
* **claude:** remove system prompt - use default community maintained prompt ([1fa7439](https://github.com/inference-gateway/sdk/commit/1fa743928dd8a4efd4b5ef1bf31b702a738a77a1))
* **claude:** standardize workflow + task-based branch prefix ([6f1616f](https://github.com/inference-gateway/sdk/commit/6f1616f9821b4e7185d68a13c8421c7af7a2c848))
* **deps:** Bump anthropics/claude-code-action  v1.0.131 -> v1.0.133 ([a4f1945](https://github.com/inference-gateway/sdk/commit/a4f19459fc8a99168edf2e3a647952e6c9ec9b35))
* **deps:** Bump anthropics/claude-code-action in the github-actions group ([#55](https://github.com/inference-gateway/sdk/issues/55)) ([d9ff0fc](https://github.com/inference-gateway/sdk/commit/d9ff0fc6ee792bac460e3bbd9ceac832904a6ce0))
* **deps:** Bump github.com/oapi-codegen/runtime in the gomod group ([#54](https://github.com/inference-gateway/sdk/issues/54)) ([26a830c](https://github.com/inference-gateway/sdk/commit/26a830c976288b7f6acd4613c9ceffd1165e5164))
* **deps:** bump golangci/golangci-lint-action in the github-actions group ([#56](https://github.com/inference-gateway/sdk/issues/56)) ([b51eaac](https://github.com/inference-gateway/sdk/commit/b51eaac60308988c50483fa3e050929e70a1cb7e))
* **deps:** bump the github-actions group with 2 updates ([#67](https://github.com/inference-gateway/sdk/issues/67)) ([23566a6](https://github.com/inference-gateway/sdk/commit/23566a6241ff1ea2d1dcf8a66e7770013a688c41))
* **deps:** Update Claude Code Action to version 1.0.131 ([df3e6b0](https://github.com/inference-gateway/sdk/commit/df3e6b02ce0e109b0e80cccaa4e5414a1d7a69ab))
* **deps:** Update claude-code-action to version 1.0.130 ([9b1dd9a](https://github.com/inference-gateway/sdk/commit/9b1dd9ae9a02816e455bc084decd120092d544cd))
* **deps:** Update golangci-lint installation to use action and add task setup ([2a9862b](https://github.com/inference-gateway/sdk/commit/2a9862b75ff53c972deb43902bd8915595c73f6d))
* **infer:** centralize infer.yml + bump infer CLI and sync .infer config ([#63](https://github.com/inference-gateway/sdk/issues/63)) ([8106a91](https://github.com/inference-gateway/sdk/commit/8106a91a41cc02ea0e4ccc614206072abe2c16f6))
* modify release workflow to use GitHub app token ([d21a8ae](https://github.com/inference-gateway/sdk/commit/d21a8aec459e29c371b0f56a6db13507a2e83ccc))

### 🔧 Miscellaneous

* **deps:** bump claude-code 2.1.148 -> 2.1.158 ([#65](https://github.com/inference-gateway/sdk/issues/65)) ([71de38e](https://github.com/inference-gateway/sdk/commit/71de38e50382afb9cbb447a229428e809e21f960))
* **deps:** bump claude-code 2.1.158 -> 2.1.161, claude-code-action v1.0.133 -> v1.0.135 ([#73](https://github.com/inference-gateway/sdk/issues/73)) ([bd7c623](https://github.com/inference-gateway/sdk/commit/bd7c6235ddaaa191555811586ab81bbab74df401))
* **deps:** bump codex 0.133.0 -> 0.135.0 ([#69](https://github.com/inference-gateway/sdk/issues/69)) ([921d2d2](https://github.com/inference-gateway/sdk/commit/921d2d27c0ab3b445f8fa8758d6cd56e48e46936))
* **deps:** Bump dev dependencies ([8bef526](https://github.com/inference-gateway/sdk/commit/8bef5261c42925cfa6d7eb50866317c50cb37c98))
* **deps:** bump infer CLI v0.117.0 -> v0.117.1, infer-action v0.9.1 -> v0.11.1 ([#64](https://github.com/inference-gateway/sdk/issues/64)) ([6be8292](https://github.com/inference-gateway/sdk/commit/6be8292a028deabf3284cbb993ebc2ec5780002e))
* **deps:** bump infer CLI v0.117.1 -> v0.119.0, infer-action v0.11.2 -> v0.11.4 ([#70](https://github.com/inference-gateway/sdk/issues/70)) ([f7223c3](https://github.com/inference-gateway/sdk/commit/f7223c35705d58e90f95be966f231771a15dcd19))
* **deps:** bump infer CLI v0.119.0 -> v0.120.0, infer-action v0.11.4 -> v0.11.6 ([#71](https://github.com/inference-gateway/sdk/issues/71)) ([e1bfba9](https://github.com/inference-gateway/sdk/commit/e1bfba91930cff5299ba20ded51691e90ab2aea2))
* **deps:** bump infer CLI v0.120.0 -> v0.120.1, infer-action v0.11.6 -> v0.11.7 ([#72](https://github.com/inference-gateway/sdk/issues/72)) ([de1ff89](https://github.com/inference-gateway/sdk/commit/de1ff89f81f134f5f929719a3e1dafdf1341decc))
* **deps:** bump infer CLI v0.120.1 -> v0.121.0 ([#74](https://github.com/inference-gateway/sdk/issues/74)) ([032bd0d](https://github.com/inference-gateway/sdk/commit/032bd0d4a00ab26f0d3a2323e9ce799caa31ed40))
* **deps:** bump infer-action v0.11.1 -> v0.11.2 ([#68](https://github.com/inference-gateway/sdk/issues/68)) ([3a9c6a6](https://github.com/inference-gateway/sdk/commit/3a9c6a6880e055196ec6823ba0be4a00f15f4b15))
* **deps:** bump infer-action v0.11.7 -> v0.12.1 ([#76](https://github.com/inference-gateway/sdk/issues/76)) ([65e438d](https://github.com/inference-gateway/sdk/commit/65e438d57a038c22e36b9c28aeef14301ee9b237))
* **deps:** Update claude-code version to 2.1.141 and infer.flake to v0.109.11 ([aa5019d](https://github.com/inference-gateway/sdk/commit/aa5019dc213d7e97ef95d47b192d489389ba0402))
* **docs:** Generate AGENTS.md file ([073632b](https://github.com/inference-gateway/sdk/commit/073632bfa2fe54490c8cedf5452337dcc4bb47f3))
* **docs:** Generate CLAUDE.md file ([3fdc2e7](https://github.com/inference-gateway/sdk/commit/3fdc2e7984cd1c7426ae5d2d13b19afaddfb065a))
* **docs:** Remove CLAUDE.md ([6364962](https://github.com/inference-gateway/sdk/commit/63649620b5ce1055411b3f972f50b004747be312))
* **flox:** add missing manifest.lock file ([c92705b](https://github.com/inference-gateway/sdk/commit/c92705b82b35e0117f9a6c517e04f8270dde088c))
* **flox:** Bump schema version ([afde3f3](https://github.com/inference-gateway/sdk/commit/afde3f3bddee655bf3365c2dd7288aff0c8972a3))
* **license:** Update license to Apache 2.0 ([033317a](https://github.com/inference-gateway/sdk/commit/033317af2034163cf183c9488f0d269f85e1a86b))
* Replace em dashes with normal dashes ([899fac7](https://github.com/inference-gateway/sdk/commit/899fac7b60fa83a3901e0357b9a812e8f5c50a3a))

## [1.16.4](https://github.com/inference-gateway/sdk/compare/v1.16.3...v1.16.4) (2026-05-19)

### 👷 CI

* **dependabot:** Add dependabot to help with dependecies upgrades ([36e91b9](https://github.com/inference-gateway/sdk/commit/36e91b9db70b720ac30f527b380d080bed4a9aa9))
* **deps:** Bump anthropics/claude-code-action ([#53](https://github.com/inference-gateway/sdk/issues/53)) ([919e4f4](https://github.com/inference-gateway/sdk/commit/919e4f448209f1e164541a38c9aa7f07668b4454))

### 🔧 Miscellaneous

* **deps:** Bump claude-code to latest and add infer CLI to flox environment ([1c2c149](https://github.com/inference-gateway/sdk/commit/1c2c1492b0f0048a6866dc37abab2d33e7781ad8))
* **deps:** Bump dev and ci dependecies to latest ([cdd9c73](https://github.com/inference-gateway/sdk/commit/cdd9c7356dc8cddfe1667364251adad57cdf31b8))

## [1.16.3](https://github.com/inference-gateway/sdk/compare/v1.16.2...v1.16.3) (2026-05-13)

### ♻️ Improvements

* Align generated code with Go naming conventions ([#52](https://github.com/inference-gateway/sdk/issues/52)) ([c491916](https://github.com/inference-gateway/sdk/commit/c491916ef835d7d848a979ff19d984083527b9e8))
* Remove unneeded copilot-instructions.md ([59b59cb](https://github.com/inference-gateway/sdk/commit/59b59cb71dc57fd46dbccf4f44a496f7cdf954e3))

### 👷 CI

* Enable display report for Claude Code action ([fa596cd](https://github.com/inference-gateway/sdk/commit/fa596cd203ad69bdc0634f96a68e1cc148eef92e))

### 📚 Documentation

* Update CLAUDE.md ([80fb78e](https://github.com/inference-gateway/sdk/commit/80fb78ed5a2e68840676588d960d2a066321ce29))

### 🔧 Miscellaneous

* Add codeowners ([db9b691](https://github.com/inference-gateway/sdk/commit/db9b6919e394c1baa8ae15585ca2a5496cefd20e))
* Add dependabot ([b9ac9ae](https://github.com/inference-gateway/sdk/commit/b9ac9ae089c01520ffa33c0f9daf44b8b9ec98e1))
* Remove devcontainers ([74ab9f6](https://github.com/inference-gateway/sdk/commit/74ab9f68e9c51e465b86e04f3497ab1ebe8ea10c))
* Remove outdated issue templates for bug reports, feature requests, and refactor requests ([e6c7ae4](https://github.com/inference-gateway/sdk/commit/e6c7ae42bbb80f6953449e04e464a4917241d42b))

### 📦 Miscellaneous

* **deps:** Bump actions/checkout from 4.2.2 to 6.0.2 ([#46](https://github.com/inference-gateway/sdk/issues/46)) ([8c9b678](https://github.com/inference-gateway/sdk/commit/8c9b67871264b08d0a8ecf746c775d7c45fba239))
* **deps:** Bump actions/setup-node from 4 to 6 ([#47](https://github.com/inference-gateway/sdk/issues/47)) ([68a3228](https://github.com/inference-gateway/sdk/commit/68a3228234c54a98b0ceb33f5e592e373c753937))
* **deps:** Bump anthropics/claude-code-action from 1.0.114 to 1.0.121 ([#48](https://github.com/inference-gateway/sdk/issues/48)) ([14e8d3d](https://github.com/inference-gateway/sdk/commit/14e8d3d39998936b59bd8b4a5c8ab47e333391b2))
* **deps:** Bump github.com/go-resty/resty/v2 from 2.16.3 to 2.17.2 ([#49](https://github.com/inference-gateway/sdk/issues/49)) ([4367734](https://github.com/inference-gateway/sdk/commit/4367734beec69a74b9fc395996ac41188f9e15d2))
* **deps:** Bump github.com/oapi-codegen/runtime from 1.1.2 to 1.4.0 ([#51](https://github.com/inference-gateway/sdk/issues/51)) ([3b9cc6c](https://github.com/inference-gateway/sdk/commit/3b9cc6cf2d2fb279d88ba8a163ae99ec1a642dfc))
* **deps:** Bump github.com/stretchr/testify from 1.10.0 to 1.11.1 ([#50](https://github.com/inference-gateway/sdk/issues/50)) ([d0211eb](https://github.com/inference-gateway/sdk/commit/d0211eb189722356fa44883acc2980d2b0c9a547))

## [1.16.2](https://github.com/inference-gateway/sdk/compare/v1.16.1...v1.16.2) (2026-05-07)

### ♻️ Improvements

* **codegen:** Add missing schemas and reconcile hand-rolled streaming types ([#44](https://github.com/inference-gateway/sdk/issues/44)) ([dfa6c08](https://github.com/inference-gateway/sdk/commit/dfa6c084073a16153e11139be1b29b2b17c8e1db)), closes [#41](https://github.com/inference-gateway/sdk/issues/41)

## [1.16.1](https://github.com/inference-gateway/sdk/compare/v1.16.0...v1.16.1) (2026-05-06)

### ♻️ Improvements

* Rename all instances of deepseek-chat to deepseek-v4-flash ([8036b02](https://github.com/inference-gateway/sdk/commit/8036b0235c903d3cc61fa1b9119f3e67e86ceca8))
* Rename all instances of deepseek-reasoner to deepseek-v4-pro ([4cb6b12](https://github.com/inference-gateway/sdk/commit/4cb6b12063a4123664860693cd3e44559b79d5e8))

### 🐛 Bug Fixes

* **ci:** Update golangci-lint installation command to specify binary output directory ([542c3e3](https://github.com/inference-gateway/sdk/commit/542c3e3f84d862014e1b7def25a70576902f8300))

### 🔧 Miscellaneous

* **deps:** Update GitHub Actions to use latest versions of checkout, setup-go, golangci-lint, and taskfile ([8c55ab7](https://github.com/inference-gateway/sdk/commit/8c55ab77c87d1f87a961080d36c0b6e68b7fcea0))
* **deps:** Update go.mod files to replace sdk version and ensure compatibility ([6c6e184](https://github.com/inference-gateway/sdk/commit/6c6e184705a3fcece67343873ada71612f7ef247))
* **deps:** Update golangci-lint installation script and version to v2.12.2 in Dockerfile and CI workflows ([dc56be3](https://github.com/inference-gateway/sdk/commit/dc56be373261e09a08c521b80853bca2af95e7a3))
* **deps:** Upgrade go to the latest version ([8f4ba65](https://github.com/inference-gateway/sdk/commit/8f4ba651dcb907e554fc8367628748cb85b09ac8))
* **openapi:** Sync vendored openapi.yaml with canonical schemas ([#43](https://github.com/inference-gateway/sdk/issues/43)) ([74402d0](https://github.com/inference-gateway/sdk/commit/74402d0b9531ccd6b9d9011edefa7ed4202c4ef6)), closes [#42](https://github.com/inference-gateway/sdk/issues/42)

## [1.16.0](https://github.com/inference-gateway/sdk/compare/v1.15.0...v1.16.0) (2026-04-28)

### ✨ Features

* Add google's thought_signature ([#35](https://github.com/inference-gateway/sdk/issues/35)) ([f123e7b](https://github.com/inference-gateway/sdk/commit/f123e7bcb63f4f73a7b5b01343a7a50b3268e2de))

## [1.15.0](https://github.com/inference-gateway/sdk/compare/v1.14.1...v1.15.0) (2026-01-24)

### ✨ Features

* Add Moonshot AI provider and sync OpenAPI schema ([#34](https://github.com/inference-gateway/sdk/issues/34)) ([91b1dd4](https://github.com/inference-gateway/sdk/commit/91b1dd4727fc27ab56f9582e1dc44859c1a40e43))

## [1.14.1](https://github.com/inference-gateway/sdk/compare/v1.14.0...v1.14.1) (2025-11-24)

### ♻️ Improvements

* **types:** Rename Message_Content to MessageContent ([#33](https://github.com/inference-gateway/sdk/issues/33)) ([5962d4c](https://github.com/inference-gateway/sdk/commit/5962d4cbb36b449d5f2cb92e143943323792e7a0)), closes [#32](https://github.com/inference-gateway/sdk/issues/32)

## [1.14.0](https://github.com/inference-gateway/sdk/compare/v1.13.0...v1.14.0) (2025-11-20)

### ⚠ BREAKING CHANGES

* **vision:** Message.Content is now MessageContent type instead of string.
Use helper functions or .FromMessageContent0() / .AsMessageContent0() methods.

Supports:
- Image URLs (https://)
- Base64-encoded images (data:image/...)
- Multiple images per message
- Image detail levels (auto, low, high)

### ✨ Features

* **provider:** Add support for Ollama Cloud ([#31](https://github.com/inference-gateway/sdk/issues/31)) ([31aaaf8](https://github.com/inference-gateway/sdk/commit/31aaaf8f339e29aeb1b4dd6aee24acfe0afc4e08))
* **vision:** Add image URL support for vision models ([#29](https://github.com/inference-gateway/sdk/issues/29)) ([b4cc118](https://github.com/inference-gateway/sdk/commit/b4cc118ab71feedcec2fd494e2d86c6575caa580)), closes [#26](https://github.com/inference-gateway/sdk/issues/26)

### 👷 CI

* Update Claude Code CI ([#27](https://github.com/inference-gateway/sdk/issues/27)) ([babade6](https://github.com/inference-gateway/sdk/commit/babade6b4d13114f9cc33c8245bcde73a7ee39a6))

### 🔧 Miscellaneous

* Delete .github/workflows/claude-code-review.yml ([#28](https://github.com/inference-gateway/sdk/issues/28)) ([9b75d24](https://github.com/inference-gateway/sdk/commit/9b75d244b8667205140c3264f0786aa1ed71806a))

## [1.13.0](https://github.com/inference-gateway/sdk/compare/v1.12.0...v1.13.0) (2025-08-28)

### ✨ Features

* **a2a:** Add functions to list and fetch A2A agents ([#25](https://github.com/inference-gateway/sdk/issues/25)) ([7af9c99](https://github.com/inference-gateway/sdk/commit/7af9c991e2349fe68d1a9cb96f7097e3cc21e411)), closes [#24](https://github.com/inference-gateway/sdk/issues/24)

## [1.12.0](https://github.com/inference-gateway/sdk/compare/v1.11.1...v1.12.0) (2025-08-22)

### ✨ Features

* **retry:** Add retry logic with exponential backoff for HTTP requests ([#20](https://github.com/inference-gateway/sdk/issues/20)) ([b4a0ddb](https://github.com/inference-gateway/sdk/commit/b4a0ddb3a00d05ec54a1d70b3184f863307b9dbb)), closes [#22](https://github.com/inference-gateway/sdk/issues/22)

## [1.12.0-rc.2](https://github.com/inference-gateway/sdk/compare/v1.12.0-rc.1...v1.12.0-rc.2) (2025-08-22)

### ✨ Features

* Add Retry Mechanism section to README and implement parseRetryAfter function with tests also for rate-limiting retries ([ce74bad](https://github.com/inference-gateway/sdk/commit/ce74badc6a0b04b66e1972157535319b3050fdca))

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
