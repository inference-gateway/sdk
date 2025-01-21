FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

ENV ZSH_CUSTOM=/home/vscode/.oh-my-zsh/custom \
    GOLANG_CI_LINT_VERSION=v1.63.4 \
    TASK_VERSION=v3.41.0

# Install common tools
RUN apt-get update && \
    # Install nodejs and npm
    curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - && \
    apt-get install -y nodejs && \
    # Install Task
    curl -s https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin ${TASK_VERSION} && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install go tools
RUN apt-get update && \
    # Install golangci-lint
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin ${GOLANG_CI_LINT_VERSION} && \
    # Clean up
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install semantic-release
RUN npm install -g semantic-release @semantic-release/changelog @semantic-release/exec @semantic-release/git @semantic-release/github conventional-changelog-conventionalcommits

# Install powerlevel10k theme
RUN git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM}/themes/powerlevel10k

USER vscode

RUN go install golang.org/x/tools/cmd/godoc@latest
