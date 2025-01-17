#!/bin/bash

set -e

git config --global --add safe.directory /workspaces/inference-gateway

# Sign commits
git config commit.gpgsign true
