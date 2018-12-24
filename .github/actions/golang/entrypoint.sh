#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail
set -e

APP_DIR="/go/src/github.com/${GITHUB_REPOSITORY}/"

mkdir -p ${APP_DIR} && cp -r ./ ${APP_DIR} && cd ${APP_DIR}

GO111MODULE=on go mod download

golangci-lint run --fast --enable-all --disable gochecknoglobals --disable gochecknoinits
