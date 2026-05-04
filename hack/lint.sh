#!/bin/bash
set -o errexit
set -x

GOLANGCI_LINT_VERSION="${GOLANGCI_LINT_VERSION:-2.9.0}"
CMD="golangci-lint run --timeout 10m0s ./..."
ENV="${ENV:-container}"
CONTAINER_ENGINE=${CONTAINER_ENGINE:-docker}

if [ "$ENV" == "container" ]; then
     ${CONTAINER_ENGINE} run --rm -v $(git rev-parse --show-toplevel):/app:Z -w /app docker.io/golangci/golangci-lint:v$GOLANGCI_LINT_VERSION $CMD
else
     curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v"$GOLANGCI_LINT_VERSION"
     $CMD
fi
