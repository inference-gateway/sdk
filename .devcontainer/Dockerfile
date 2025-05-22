FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

ENV ZSH_CUSTOM=/home/vscode/.oh-my-zsh/custom \
    GOLANG_CI_LINT_VERSION=v1.63.4 \
    TASK_VERSION=v3.41.0

# Install common tools
RUN apt-get update && \
    # Install GitHub CLI
    curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg && \
    chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null && \
    apt-get update && \
    apt-get install -y gh && \
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

# Install claude-code
RUN npm install -g @anthropic-ai/claude-code

# Install powerlevel10k theme
RUN git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM}/themes/powerlevel10k

USER vscode

RUN go install golang.org/x/tools/cmd/godoc@latest

# Use Powerlevel10k theme
RUN git clone --depth=1 https://github.com/romkatv/powerlevel10k.git /home/vscode/.powerlevel10k && \
    echo 'source /home/vscode/.powerlevel10k/powerlevel10k.zsh-theme' >> /home/vscode/.zshrc && \
    echo 'POWERLEVEL9K_DISABLE_CONFIGURATION_WIZARD=true' >> /home/vscode/.zshrc

RUN echo "source <(task --completion zsh)" >> /home/vscode/.zshrc 
