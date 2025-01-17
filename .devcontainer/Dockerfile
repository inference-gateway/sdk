FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

ENV ZSH_CUSTOM=/home/vscode/.oh-my-zsh/custom \
    GORELEASER_VERSION=v2.5.1 \
    GOLANG_CI_LINT_VERSION=v1.63.4

RUN apt-get update && \
    # Determine architecture
    architecture=$(uname -m); \
    architecture_pretty="arm64"; \
    if [ "$architecture" = "x86_64" ]; then \
    architecture_pretty="amd64"; \
    fi && \
    echo "Installing depenenices for ${architecture}" && \
    # Install powerlevel10k theme
    git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM}/themes/powerlevel10k && \
    # Install GoReleaser
    curl -sSL https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_${architecture_pretty}.tar.gz | tar -xzv -C /usr/local/bin goreleaser && \
    # Install golangci-lint
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin ${GOLANG_CI_LINT_VERSION} && \
    # Clean up
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

USER vscode
