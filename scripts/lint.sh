#!/bin/bash

# set -e: Exit immediately if a command exits with a non-zero status.
set -e
readonly service="$1"

# Navigate to the appropriate directory based on the provided service argument.
#
# If the service is "pkg", we navigate to the internal/pkg directory. Otherwise, check if the service variable is not empty and navigate to the corresponding internal/services/{service} directory.
if [ "$service" = "pkg" ]; then
    cd "./internal/pkg"
elif [ -n "$service" ]; then
    cd "./internal/services/$service"
fi

# https://golangci-lint.run/docs/welcome/quick-start/#linting
golangci-lint run ./...