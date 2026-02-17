#!/bin/bash

# set -e: Exit immediately if a command exits with a non-zero status.
set -e
readonly service="$1"

cd "./internal/services/$service" && go run "./cmd/app/main.go"